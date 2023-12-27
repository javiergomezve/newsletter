"use client";

import { Fragment } from "react";
import Breadcrumb from "@/components/Breadcrumbs/Breadcrumb";
import { useGetNewsLettersQuery } from "@/redux/services/newslettersApi";

const NewslettersPage = () => {

    const {
        data: newsletters,
    } = useGetNewsLettersQuery(null);

    return (
        <Fragment>
            <Breadcrumb pageName="Newsletters" />

            <div className="flex flex-col gap-10">
                <div
                    className="rounded-sm border border-stroke bg-white px-5 pt-6 pb-2.5 shadow-default dark:border-strokedark dark:bg-boxdark sm:px-7.5 xl:pb-1"
                >
                    <div className="flex flex-col">
                        <div className="grid grid-cols-3 rounded-sm bg-gray-2 dark:bg-meta-4">
                            <div className="p-2.5 xl:p-5">
                                <h5 className="text-sm font-medium uppercase xsm:text-base">
                                    Subject
                                </h5>
                            </div>
                            <div className="p-2.5 text-center xl:p-5">
                                <h5 className="text-sm font-medium uppercase xsm:text-base">
                                    Send at
                                </h5>
                            </div>
                            <div className="p-2.5 text-center xl:p-5">
                                <h5 className="text-sm font-medium uppercase xsm:text-base">
                                    Actions
                                </h5>
                            </div>
                        </div>

                        {newsletters?.map((newsletter, key) => {
                            return (
                                <div
                                    className={`grid grid-cols-3 ${
                                        key === newsletters?.length - 1
                                            ? ""
                                            : "border-b border-stroke dark:border-strokedark"
                                    }`}
                                    key={newsletter.id}
                                >
                                    <div className="flex items-center gap-3 p-2.5 xl:p-5">
                                        <p className="hidden text-black dark:text-white sm:block">
                                            {newsletter.subject}
                                        </p>
                                    </div>

                                    <div className="flex items-center justify-center p-2.5 xl:p-5">
                                        <p className="text-black dark:text-white">
                                            {formatDate(new Date(newsletter.send_at))}
                                        </p>
                                    </div>

                                    <div className="flex items-center justify-center p-2.5 xl:p-5">
                                        <button
                                            onClick={() => {
                                                // window.open(media.location, '_blank');
                                            }}
                                            // disabled={isDeleting || isCreating || isUpdating}
                                        >
                                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                                                 strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                                                <path strokeLinecap="round" strokeLinejoin="round"
                                                      d="M3 16.5v2.25A2.25 2.25 0 0 0 5.25 21h13.5A2.25 2.25 0 0 0 21 18.75V16.5M16.5 12 12 16.5m0 0L7.5 12m4.5 4.5V3" />
                                            </svg>
                                        </button>
                                    </div>
                                </div>
                            )
                        })}
                    </div>
                </div>
            </div>
        </Fragment>
    )
};

export default NewslettersPage;

function formatDate(date: Date): string {
    const formatOptions: Intl.DateTimeFormatOptions = { month: '2-digit', day: '2-digit', year: 'numeric' };
    return date.toLocaleDateString(undefined, formatOptions);
}