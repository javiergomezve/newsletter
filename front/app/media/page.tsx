"use client";

import { ChangeEvent, Fragment, useCallback, useEffect, useState } from "react";
import Breadcrumb from "@/components/Breadcrumbs/Breadcrumb";
import {
    ErrorResponse,
    Media,
    useCreateMediaMutation,
    useDeleteMediaMutation,
    useGetMediasQuery,
    useUpdateMediaMutation
} from "@/redux/services/mediaApi";
import Loader from "@/components/common/Loader";
import Image from "next/image";
import Modal from "@/components/Modal";
import Alert, { AlertProps } from "@/components/Alert";

const MediasPage = () => {

    const {
        data: medias, error, isLoading, isFetching, refetch: refetchMedias
    } = useGetMediasQuery(null);
    const [deleteMedia, { isLoading: isDeleting }] = useDeleteMediaMutation();
    const [createMedia, { isLoading: isCreating }] = useCreateMediaMutation();
    const [updateMedia, { isLoading: isUpdating }] = useUpdateMediaMutation();

    const [selectedFile, setSelectedFile] = useState<File | null>(null);
    const [mediaToDelete, setMediaToDelete] = useState<Media | null>(null);
    const [mediaToEdit, setMediaToEdit] = useState<Media | null>(null);

    const [showAlert, setShowAlert] = useState<AlertProps | null>(null);

    const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
        const file = e.target.files?.[0];

        if (file) setSelectedFile(file);
    };

    const handleCreateMedia = useCallback(async () => {
        if (!selectedFile) return;

        const formData = new FormData();
        formData.append("media[]", selectedFile);

        const response = await createMedia(formData);
        if ("error" in response) {
            const error = response.error as ErrorResponse;

            setShowAlert({
                type: "danger",
                title: error.data.message,
                message: ""
            });
            return;
        }

        await refetchMedias();

        setSelectedFile(null);

        setShowAlert({
            type: "success",
            title: "Media created successfully",
            message: ""
        });
    }, [selectedFile]);

    const handleDeleteMedia = useCallback(async () => {
        setMediaToDelete(null);

        const response = await deleteMedia({ id: mediaToDelete!.id });
        if ("error" in response) {
            const error = response.error as ErrorResponse;

            setShowAlert({
                type: "danger",
                title: error.data.message,
                message: ""
            });
            return;
        }

        await refetchMedias();

        setShowAlert({
            type: "success",
            title: "Media deleted successfully",
            message: ""
        });
    }, [mediaToDelete]);

    const handleEditMedia = useCallback(async () => {
        if (!selectedFile) return;

        const formData = new FormData();
        formData.append("media[]", selectedFile);
        const response = await updateMedia({ id: mediaToEdit!.id, formData });

        if ("error" in response) {
            const error = response.error as ErrorResponse;

            setShowAlert({
                type: "danger",
                title: error.data.message,
                message: ""
            });
            return;
        }

        await refetchMedias();

        setSelectedFile(null);
        setMediaToEdit(null);
    }, [mediaToEdit, selectedFile]);

    const handleSubmit = useCallback((event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        if (mediaToEdit) {
            handleEditMedia();
        } else {
            handleCreateMedia();
        }
    }, [selectedFile, mediaToEdit]);

    useEffect(() => {
        if (showAlert) {
            setTimeout(() => {
                setShowAlert(null);
            }, 3000);
        }
    }, [showAlert]);

    if (isLoading) {
        return <Loader />;
    }

    if (error) {
        return <p>error</p>;
    }

    return (
        <Fragment>
            <Modal
                isOpen={!!mediaToDelete}
                title={"Delete media"}
                body={<p>
                    Are yo sure you want to delete {mediaToDelete?.file_name} ?
                </p>}
                onClose={() => {
                    setMediaToDelete(null);
                }}
                onSubmit={handleDeleteMedia}
            />

            <Breadcrumb pageName="Media" />

            {
                showAlert && (
                    <div className="mb-2">
                        <Alert type={showAlert.type} title={showAlert.title} message={showAlert.message} />
                    </div>
                )
            }

            <div className="grid grid-cols-5 gap-8">
                <div className="col-span-5 xl:col-span-3">
                    <div className="flex flex-col gap-10">
                        <div
                            className="rounded-sm border border-stroke bg-white px-5 pt-6 pb-2.5 shadow-default dark:border-strokedark dark:bg-boxdark sm:px-7.5 xl:pb-1"
                        >
                            <div className="flex flex-col">
                                <div className="grid grid-cols-3 rounded-sm bg-gray-2 dark:bg-meta-4">
                                    <div className="p-2.5 xl:p-5">
                                        <h5 className="text-sm font-medium uppercase xsm:text-base">
                                            Name
                                        </h5>
                                    </div>
                                    <div className="p-2.5 text-center xl:p-5">
                                        <h5 className="text-sm font-medium uppercase xsm:text-base">
                                            Type
                                        </h5>
                                    </div>
                                    <div className="p-2.5 text-center xl:p-5">
                                        <h5 className="text-sm font-medium uppercase xsm:text-base">
                                            Actions
                                        </h5>
                                    </div>
                                </div>

                                {medias?.map((media, key) => (
                                    <div
                                        className={`grid grid-cols-3 ${
                                            key === medias?.length - 1
                                                ? ""
                                                : "border-b border-stroke dark:border-strokedark"
                                        }`}
                                        key={key}
                                    >
                                        <div className="flex items-center gap-3 p-2.5 xl:p-5">
                                            <p className="hidden text-black dark:text-white sm:block">
                                                {media.file_name}
                                            </p>
                                        </div>

                                        <div className="flex items-center justify-center p-2.5 xl:p-5">
                                            <p className="text-black dark:text-white">{media.content_type}</p>
                                        </div>

                                        <div className="flex items-center justify-center p-2.5 xl:p-5">
                                            <button
                                                onClick={() => {
                                                    window.open(media.location, "_blank");
                                                }}
                                                disabled={isDeleting || isCreating || isUpdating}
                                            >
                                                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                                                     strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                                                    <path strokeLinecap="round" strokeLinejoin="round"
                                                          d="M3 16.5v2.25A2.25 2.25 0 0 0 5.25 21h13.5A2.25 2.25 0 0 0 21 18.75V16.5M16.5 12 12 16.5m0 0L7.5 12m4.5 4.5V3" />
                                                </svg>
                                            </button>

                                            <button
                                                onClick={() => {
                                                    setMediaToEdit(media);
                                                }}
                                                disabled={isDeleting || isCreating || isUpdating}
                                            >
                                                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                                                     strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                                                    <path strokeLinecap="round" strokeLinejoin="round"
                                                          d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L6.832 19.82a4.5 4.5 0 0 1-1.897 1.13l-2.685.8.8-2.685a4.5 4.5 0 0 1 1.13-1.897L16.863 4.487Zm0 0L19.5 7.125" />
                                                </svg>
                                            </button>

                                            <button
                                                onClick={() => {
                                                    setMediaToDelete(media);
                                                }}
                                                disabled={isDeleting || isCreating || isUpdating}
                                            >
                                                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                                                     strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                                                    <path strokeLinecap="round" strokeLinejoin="round"
                                                          d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0" />
                                                </svg>
                                            </button>
                                        </div>
                                    </div>
                                ))}
                            </div>
                        </div>
                    </div>
                </div>

                <div className="col-span-5 xl:col-span-2">
                    <div
                        className="rounded-sm border border-stroke bg-white shadow-default dark:border-strokedark dark:bg-boxdark"
                    >
                        <div className="border-b border-stroke py-4 px-7 dark:border-strokedark">
                            <h3 className="font-medium text-black dark:text-white">
                                {mediaToEdit
                                    ? "Edit " + mediaToEdit.file_name
                                    : "Upload a new media"
                                }
                            </h3>
                        </div>
                        <div className="p-7">
                            <form onSubmit={handleSubmit}>
                                <div
                                    id="FileUpload"
                                    className="relative mb-5.5 block w-full cursor-pointer appearance-none rounded border-2 border-dashed border-primary bg-gray py-4 px-4 dark:bg-meta-4 sm:py-7.5"
                                >
                                    <input
                                        type="file"
                                        accept=".pdf, image/png"
                                        onChange={handleFileChange}
                                        className="absolute inset-0 z-50 m-0 h-full w-full cursor-pointer p-0 opacity-0 outline-none"
                                    />
                                    <div className="flex flex-col items-center justify-center space-y-3">
                                        {(!selectedFile && !mediaToEdit) && (
                                            <Fragment>
                                            <span
                                                className="flex h-10 w-10 items-center justify-center rounded-full border border-stroke bg-white dark:border-strokedark dark:bg-boxdark">
                        <svg
                            width="16"
                            height="16"
                            viewBox="0 0 16 16"
                            fill="none"
                            xmlns="http://www.w3.org/2000/svg"
                        >
                          <path
                              fillRule="evenodd"
                              clipRule="evenodd"
                              d="M1.99967 9.33337C2.36786 9.33337 2.66634 9.63185 2.66634 10V12.6667C2.66634 12.8435 2.73658 13.0131 2.8616 13.1381C2.98663 13.2631 3.1562 13.3334 3.33301 13.3334H12.6663C12.8431 13.3334 13.0127 13.2631 13.1377 13.1381C13.2628 13.0131 13.333 12.8435 13.333 12.6667V10C13.333 9.63185 13.6315 9.33337 13.9997 9.33337C14.3679 9.33337 14.6663 9.63185 14.6663 10V12.6667C14.6663 13.1971 14.4556 13.7058 14.0806 14.0809C13.7055 14.456 13.1968 14.6667 12.6663 14.6667H3.33301C2.80257 14.6667 2.29387 14.456 1.91879 14.0809C1.54372 13.7058 1.33301 13.1971 1.33301 12.6667V10C1.33301 9.63185 1.63148 9.33337 1.99967 9.33337Z"
                              fill="#3C50E0"
                          />
                          <path
                              fillRule="evenodd"
                              clipRule="evenodd"
                              d="M7.5286 1.52864C7.78894 1.26829 8.21106 1.26829 8.4714 1.52864L11.8047 4.86197C12.0651 5.12232 12.0651 5.54443 11.8047 5.80478C11.5444 6.06513 11.1223 6.06513 10.8619 5.80478L8 2.94285L5.13807 5.80478C4.87772 6.06513 4.45561 6.06513 4.19526 5.80478C3.93491 5.54443 3.93491 5.12232 4.19526 4.86197L7.5286 1.52864Z"
                              fill="#3C50E0"
                          />
                          <path
                              fillRule="evenodd"
                              clipRule="evenodd"
                              d="M7.99967 1.33337C8.36786 1.33337 8.66634 1.63185 8.66634 2.00004V10C8.66634 10.3682 8.36786 10.6667 7.99967 10.6667C7.63148 10.6667 7.33301 10.3682 7.33301 10V2.00004C7.33301 1.63185 7.63148 1.33337 7.99967 1.33337Z"
                              fill="#3C50E0"
                          />
                        </svg>
                      </span>
                                                <p>
                                                    <span className="text-primary">Click to upload</span> or
                                                    drag and drop
                                                </p>
                                                <p className="mt-1.5">PDF or PNG</p>
                                                <p>(max 8MB)</p>
                                            </Fragment>
                                        )}

                                        {selectedFile && (
                                            <p>{selectedFile.name}</p>
                                        )}

                                        {(mediaToEdit && !selectedFile) && (
                                            <p>{mediaToEdit.file_name}</p>
                                        )}
                                    </div>
                                </div>

                                <div className="flex justify-end gap-4.5">
                                    <button
                                        className="flex justify-center rounded bg-primary py-2 px-6 font-medium text-gray hover:bg-opacity-95"
                                        type="submit"
                                        disabled={isDeleting || isCreating || isUpdating}
                                    >
                                        Save
                                    </button>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </Fragment>
    );
};

export default MediasPage;
