"use client";

import Breadcrumb from "@/components/Breadcrumbs/Breadcrumb";
import { useGetNewsLettersQuery } from "@/redux/services/newslettersApi";
import AdminLayout from "@/components/AdminLayout";

const NewslettersPage = () => {

    const {
        data: newsletters,
    } = useGetNewsLettersQuery(null);

    return (
        <AdminLayout>
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
                                    Data
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

                                    <div className="flex flex-col items-center justify-center p-2.5 xl:p-5">
                                        <p className="text-black dark:text-white">
                                            Recipients: {newsletter.recipients?.length}
                                        </p>

                                        <p className="text-black dark:text-white">
                                            Attachments: {newsletter.attachments?.length}
                                        </p>
                                    </div>
                                </div>
                            )
                        })}
                    </div>
                </div>
            </div>
        </AdminLayout>
    )
};

export default NewslettersPage;

function formatDate(date: Date): string {
    const formatOptions: Intl.DateTimeFormatOptions = { month: '2-digit', day: '2-digit', year: 'numeric' };
    return date.toLocaleDateString(undefined, formatOptions);
}