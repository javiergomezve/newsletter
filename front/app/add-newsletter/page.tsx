"use client";

import { useEffect, useState } from "react";
import "react-quill/dist/quill.snow.css";
import { Newsletter, useCreateNewsletterMutation } from "@/redux/services/newslettersApi";
import { useGetSubscribersQuery } from "@/redux/services/recipientsApi";
import { ErrorResponse, useGetMediasQuery } from "@/redux/services/mediaApi";
import Breadcrumb from "@/components/Breadcrumbs/Breadcrumb";
import Alert, { AlertProps } from "@/components/Alert";
import Checkbox from "@/components/Checkbox";
import Loader from "@/components/common/Loader";
import dynamic from "next/dynamic";
import AdminLayout from "@/components/AdminLayout";

const ReactQuill = dynamic(
    () => import("react-quill"),
    { ssr: false }
);

const defaultNewsletter = {
    subject: "",
    content: "",
    send_at: "",
    recipients: {},
    attachments: {}
};

const AddNewslettersPage = () => {
    const [
        createNewsletter,
        {
            isLoading: creatingNewsletter
        }
    ] = useCreateNewsletterMutation();

    const {
        data: recipients,
        isLoading: loadingRecipients,
        error: errorLoadingRecipients
    } = useGetSubscribersQuery(null);

    const {
        data: medias,
        isLoading: loadingMedias,
        error: errorLoadingMedias
    } = useGetMediasQuery(null);

    const [newsletter, setNewsletter] = useState<Newsletter>(defaultNewsletter);

    const [showAlert, setShowAlert] = useState<AlertProps | null>(null);

    const handleChangeSubject = (subject: string) => {
        setNewsletter({
            ...newsletter,
            subject
        });
    };

    const handleChangeContent = (content: string) => {
        setNewsletter({
            ...newsletter,
            content
        });
    };

    const toggleRecipient = (id: string) => {
        const newRecipients = { ...newsletter.recipients };

        if (id in newRecipients) {
            delete newRecipients[id];
        } else {
            newRecipients[id] = true;
        }

        setNewsletter({
            ...newsletter,
            recipients: newRecipients
        });
    };

    const handleSelectAllRecipients = () => {
        let selectedRecipient: { [recipientId: string]: boolean } = {};

        recipients?.map(recipient => {
            // @ts-ignore
            selectedRecipient[recipient?.id] = true;
        });

        setNewsletter({
            ...newsletter,
            recipients: selectedRecipient
        });
    };

    const toggleAttachments = (id: string) => {
        const newAttachments = { ...newsletter.attachments };

        if (id in newAttachments) {
            delete newAttachments[id];
        } else {
            newAttachments[id] = true;
        }

        setNewsletter({
            ...newsletter,
            attachments: newAttachments
        });
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        if (
            !newsletter.subject ||
            !newsletter.content ||
            !(newsletter.recipients && Object.keys(newsletter.recipients).length > 0)
        ) return;

        const response = await createNewsletter(newsletter);
        if ("error" in response) {
            const error = response.error as ErrorResponse;

            setShowAlert({
                type: "danger",
                title: error.data.message,
                message: ""
            });
            return;
        }

        setShowAlert({ type: "success", title: "Newsletter created successfully", message: "" });

        setNewsletter(defaultNewsletter);
    };

    useEffect(() => {
        if (showAlert) {
            setTimeout(() => {
                setShowAlert(null);
            }, 3000);
        }
    }, [showAlert]);

    if (loadingRecipients || loadingMedias) {
        return <Loader />;
    }

    if (errorLoadingRecipients || errorLoadingMedias) {
        return <p>error</p>;
    }

    return (
        <AdminLayout>
            <Breadcrumb pageName="Add Newsletter" />

            {
                showAlert && (
                    <div className="mb-2">
                        <Alert type={showAlert.type} title={showAlert.title} message={showAlert.message} />
                    </div>
                )
            }

            <div
                className="rounded-sm border border-stroke bg-white shadow-default dark:border-strokedark dark:bg-boxdark"
            >
                <div className="border-b border-stroke py-4 px-6.5 dark:border-strokedark">
                    <h3 className="font-medium text-black dark:text-white">
                        New newsletter
                    </h3>
                </div>

                <form onSubmit={handleSubmit}>
                    <div className="p-6.5">
                        <div className="mb-4.5">
                            <label className="mb-2.5 block text-black dark:text-white">
                                Subject
                            </label>
                            <input
                                type="text"
                                placeholder="Select subject"
                                className="w-full rounded border-[1.5px] border-stroke bg-transparent py-3 px-5 font-medium outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-whiter dark:border-form-strokedark dark:bg-form-input dark:focus:border-primary"
                                value={newsletter.subject}
                                onChange={(event) => handleChangeSubject(event.target.value)}
                            />
                        </div>

                        <div className="mb-4.5">
                            <div className="flex justify-between">
                                <label className="mb-2.5 block text-black dark:text-white">
                                    Recipients
                                </label>

                                <button
                                    className="flex  justify-center rounded bg-primary p-3 font-medium text-gray"
                                    type="button"
                                    disabled={creatingNewsletter}
                                    onClick={handleSelectAllRecipients}
                                >
                                    check all
                                </button>
                            </div>

                            <div className="flex">
                                {recipients?.map((recipient) => {
                                    let isChecked = false;
                                    if (recipient.id && newsletter.recipients) {
                                        isChecked = recipient.id in newsletter.recipients;
                                    }

                                    return (
                                        <Checkbox
                                            key={recipient.id}
                                            containerClassName="mr-2"
                                            label={recipient.email}
                                            isChecked={isChecked}
                                            onChange={() => {
                                                if (recipient.id) {
                                                    toggleRecipient(recipient.id);
                                                }
                                            }}
                                        />
                                    );
                                })}
                            </div>
                        </div>

                        <div className="mb-6">
                            <label className="mb-2.5 block text-black dark:text-white">
                                Message
                            </label>

                            <ReactQuill
                                theme="snow"
                                modules={{
                                    toolbar: [
                                        [{ "header": [1, 2, false] }],
                                        ["bold", "italic", "underline", "strike", "blockquote"],
                                        [{ "list": "ordered" }, { "list": "bullet" }, { "indent": "-1" }, { "indent": "+1" }],
                                        ["link", "image"],
                                        ["clean"]
                                    ]
                                }}
                                value={newsletter.content}
                                onChange={(value) => handleChangeContent(value)}
                            />
                        </div>

                        <div className="mb-4.5">
                            <label className="mb-2.5 block text-black dark:text-white">
                                Attachments
                            </label>

                            <div className="flex">
                                {medias?.map((media) => {
                                    let isChecked = false;
                                    if (media.id && newsletter.attachments) {
                                        isChecked = media.id in newsletter.attachments;
                                    }

                                    return (
                                        <Checkbox
                                            key={media.id}
                                            containerClassName="mr-2"
                                            label={media.file_name}
                                            isChecked={isChecked}
                                            onChange={() => {
                                                if (media.id) {
                                                    toggleAttachments(media.id);
                                                }
                                            }}
                                        />
                                    );
                                })}
                            </div>
                        </div>

                        <button
                            className="flex w-full justify-center rounded bg-primary p-3 font-medium text-gray"
                            disabled={creatingNewsletter}
                        >
                            Send Newsletter
                        </button>
                    </div>
                </form>
            </div>
        </AdminLayout>
    );
};

export default AddNewslettersPage;
