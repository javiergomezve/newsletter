import PublicLayout from "@/components/PublicLayout";
import UnsubscribeForm from "@/app/unsubscribe/UnsubscribeForm";

interface UnsubscribePageProps {
    params: {
        slug: string;
    };
    searchParams: {
        [key: string]: string | undefined;
    };
}

const UnsubscribePage = ({ params, searchParams }: UnsubscribePageProps) => {

    return (
        <PublicLayout>
            <div className="flex flex-col gap-9 justify-center items-center min-h-screen">
                <div
                    className="rounded-sm border border-stroke bg-white shadow-default dark:border-strokedark dark:bg-boxdark w-[30%]"
                >
                    <div className="border-b border-stroke py-4 px-6.5 dark:border-strokedark">
                        <h3 className="font-medium text-black dark:text-white">
                            Unsubscribe
                        </h3>
                    </div>

                    <UnsubscribeForm email={searchParams.email} />
                </div>
            </div>
        </PublicLayout>
    );
};

export default UnsubscribePage;