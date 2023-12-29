"use client";

import "./globals.css";
import "./data-tables-css.css";
import "./satoshi.css";

import { Providers } from "@/redux/provider";

const RootLayout = ({ children }: { children: React.ReactNode; }) => {
    return (
        <html lang="en">
        <body suppressHydrationWarning={true}>
        <Providers>
            {children}
        </Providers>
        </body>
        </html>
    );
};

export default RootLayout;