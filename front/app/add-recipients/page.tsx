"use client";

import { Fragment, useCallback, useEffect, useState } from "react";
import { v4 as uuidv4 } from "uuid";

import Breadcrumb from "@/components/Breadcrumbs/Breadcrumb";
import { Recipient, useCreateRecipientsMutation } from "@/redux/services/recipientsApi";
import Alert, { AlertProps } from "@/components/Alert";
import { ErrorResponse } from "@/redux/services/mediaApi";

const AddRecipientsPage = () => {

    const [createRecipients, { isLoading: creatingRecipients }] = useCreateRecipientsMutation();

    const [recipients, setRecipients] = useState<Recipient[]>([]);
    const [showAlert, setShowAlert] = useState<AlertProps | null>(null);

    const handleAddRecipient = () => {
        const newRecipient: Recipient = {
            id: uuidv4(),
            full_name: "",
            email: ""
        };

        setRecipients([...recipients, newRecipient]);
    };

    const handleSetFullName = (id: string | undefined, fullName: string) => {
        if (!id) return;

        const index = recipients.findIndex(r => r.id === id);
        if (index > -1) {
            const newRecipients = [...recipients];
            newRecipients[index].full_name = fullName;
            setRecipients(newRecipients);
        }
    };

    const handleSetEmail = (id: string | undefined, email: string) => {
        if (!id) return;

        const index = recipients.findIndex(r => r.id === id);
        if (index > -1) {
            const newRecipients = [...recipients];
            newRecipients[index].email = email;
            setRecipients(newRecipients);
        }
    };

    const handleSubmit = useCallback(async () => {
        const recipientsToSend: Recipient[] = recipients
            .filter(recipient => recipient?.full_name && recipient?.email)
            .map(r => ({ full_name: r.full_name, email: r.email }));

        const response = await createRecipients(recipientsToSend);
        if ("error" in response) {
            const error = response.error as ErrorResponse;
            setShowAlert({
                type: "danger",
                title: error.data.message,
                message: ""
            });
            return;
        }

        setShowAlert({
            type: "success",
            title: "Recipients created successfully",
            message: ""
        });

        setRecipients([{
            id: uuidv4(),
            full_name: "",
            email: ""
        }]);
    }, [recipients]);

    useEffect(() => {
        handleAddRecipient();
    }, []);

    useEffect(() => {
        if (showAlert) {
            setTimeout(() => {
                setShowAlert(null);
            }, 3000);
        }
    }, [showAlert]);

    return (
        <Fragment>
            <Breadcrumb pageName="Add Recipients" />

            {
                showAlert && (
                    <div className="mb-2">
                        <Alert type={showAlert.type} title={showAlert.title} message={showAlert.message} />
                    </div>
                )
            }

            <div className="col-span-5 xl:col-span-3">
                <div className="flex flex-col gap-10">
                    <div className="flex">
                        <button
                            onClick={handleAddRecipient}
                            disabled={creatingRecipients}
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                                 strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                                <path strokeLinecap="round" strokeLinejoin="round"
                                      d="M12 9v6m3-3H9m12 0a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
                            </svg>
                        </button>

                        <p>Add row</p>
                    </div>

                    <div
                        className="rounded-sm border border-stroke bg-white px-5 pt-6 pb-2.5 shadow-default dark:border-strokedark dark:bg-boxdark sm:px-7.5 xl:pb-1"
                    >
                        <div className="flex flex-col">
                            <div className="grid grid-cols-2 rounded-sm bg-gray-2 dark:bg-meta-4">
                                <div className="p-2.5 xl:p-5">
                                    <h5 className="text-sm font-medium uppercase xsm:text-base">
                                        Full Name
                                    </h5>
                                </div>
                                <div className="p-2.5 text-center xl:p-5">
                                    <h5 className="text-sm font-medium uppercase xsm:text-base">
                                        Email
                                    </h5>
                                </div>
                            </div>

                            {recipients.map((recipient, index) => (
                                <div
                                    className={`grid grid-cols-2 ${index === recipients?.length - 1
                                        ? ""
                                        : "border-b border-stroke dark:border-strokedark"
                                    }`}
                                    key={recipient.id}
                                >
                                    <div className="flex items-center gap-3 p-2.5 xl:p-5">
                                        <input
                                            type="text"
                                            placeholder="Jhon Doe"
                                            className="w-full rounded-lg border-[1.5px] border-stroke bg-transparent py-3 px-5 font-medium outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-whiter dark:border-form-strokedark dark:bg-form-input dark:focus:border-primary"
                                            value={recipient.full_name}
                                            onChange={e => handleSetFullName(recipient.id, e.target.value)}
                                        />
                                    </div>

                                    <div className="flex items-center justify-center p-2.5 xl:p-5">
                                        <input
                                            type="text"
                                            placeholder="example@gmail.com"
                                            className="w-full rounded-lg border-[1.5px] border-stroke bg-transparent py-3 px-5 font-medium outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-whiter dark:border-form-strokedark dark:bg-form-input dark:focus:border-primary"
                                            value={recipient.email}
                                            onChange={e => handleSetEmail(recipient.id, e.target.value)}
                                        />
                                    </div>
                                </div>
                            ))}

                            <button
                                className="flex w-full justify-center rounded bg-primary p-3 font-medium text-gray mb-2"
                                onClick={handleSubmit}
                                disabled={creatingRecipients}
                            >
                                Save
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </Fragment>
    );
};

export default AddRecipientsPage;