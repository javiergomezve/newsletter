"use client";

import { useGetRecipientsQuery } from "@/redux/services/recipientsApi";
import Loader from "@/components/common/Loader";
import { Fragment } from "react";
import Breadcrumb from "@/components/Breadcrumbs/Breadcrumb";

const RecipientsPage = () => {

    const {
        data: recipients,
        isLoading,
        error
    } = useGetRecipientsQuery(null);

    if (isLoading) {
        return <Loader />;
    }

    if (error) {
        return <p>error</p>;
    }

    return (
        <Fragment>
            <Breadcrumb pageName="Recipients" />

            <div className="col-span-5 xl:col-span-3">
                <div className="flex flex-col gap-10">
                    <div
                        className="rounded-sm border border-stroke bg-white px-5 pt-6 pb-2.5 shadow-default dark:border-strokedark dark:bg-boxdark sm:px-7.5 xl:pb-1"
                    >
                        <div className="flex flex-col">
                            <div className="grid grid-cols-3 rounded-sm bg-gray-2 dark:bg-meta-4">
                                <div className="p-2.5 xl:p-5">
                                    <h5 className="text-sm font-medium uppercase xsm:text-base">
                                        Full name
                                    </h5>
                                </div>
                                <div className="p-2.5 text-center xl:p-5">
                                    <h5 className="text-sm font-medium uppercase xsm:text-base">
                                        Email
                                    </h5>
                                </div>
                                <div className="p-2.5 text-center xl:p-5">
                                    <h5 className="text-sm font-medium uppercase xsm:text-base">
                                        Status
                                    </h5>
                                </div>
                            </div>

                            {recipients?.map((recipient, key) => (
                                <div
                                    className={`grid grid-cols-3 ${
                                        key === recipients?.length - 1
                                            ? ""
                                            : "border-b border-stroke dark:border-strokedark"
                                    }`}
                                    key={key}
                                >
                                    <div className="flex items-center gap-3 p-2.5 xl:p-5">
                                        <p className="hidden text-black dark:text-white sm:block">
                                            {recipient.full_name}
                                        </p>
                                    </div>

                                    <div className="flex items-center justify-center p-2.5 xl:p-5">
                                        <p className="text-black dark:text-white">
                                            {recipient.email}
                                        </p>
                                    </div>

                                    <div className="flex items-center justify-center p-2.5 xl:p-5">
                                        <p className="text-black dark:text-white">
                                            {recipient?.status}
                                        </p>
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>
                </div>
            </div>
        </Fragment>
    );
};

export default RecipientsPage;