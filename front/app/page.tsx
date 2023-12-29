import ECommerce from "@/components/Dashboard/E-commerce";
import { Metadata } from "next";
import AdminLayout from "@/components/AdminLayout";

export const metadata: Metadata = {
    title: "TailAdmin | Next.js E-commerce Dashboard Template",
    description: "This is Home Blog page for TailAdmin Next.js"
};

const Home = () => {
    return (
        <AdminLayout>
            <ECommerce />
        </AdminLayout>
    );
};

export default Home;