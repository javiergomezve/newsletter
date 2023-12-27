import React from "react";

export interface AlertProps {
    type: "warning" | "success" | "danger";
    title: string;
    message: string;
}

const Alert: React.FC<AlertProps> = ({ type, title, message }) => {
    const alertColors = {
        warning: "bg-warning bg-opacity-[15%] border-warning",
        success: "bg-[#34D399] bg-opacity-[15%] border-[#34D399]",
        danger: "bg-[#F87171] bg-opacity-[15%] border-[#F87171]"
    };

    const iconColors = {
        warning: "bg-warning",
        success: "bg-[#34D399]",
        danger: "bg-[#F87171]"
    };

    return (
        <div
            className={`flex w-full border-l-6 ${alertColors[type]} px-7 py-8 shadow-md dark:bg-[#1B1B24] dark:bg-opacity-30 md:p-9`}>
            <div
                className={`mr-5 flex h-9 w-full max-w-[36px] items-center justify-center rounded-lg ${iconColors[type]}`}>
                {/* Puedes ajustar los íconos según tu necesidad */}
                {type === "warning" && (
                    <svg width="19" height="16" viewBox="0 0 19 16" fill="none" xmlns="http://www.w3.org/2000/svg">
                    </svg>
                )}

                {type === "success" && (
                    <svg width="16" height="12" viewBox="0 0 16 12" fill="none" xmlns="http://www.w3.org/2000/svg">
                    </svg>
                )}

                {type === "danger" && (
                    <svg width="13" height="13" viewBox="0 0 13 13" fill="none" xmlns="http://www.w3.org/2000/svg">
                    </svg>
                )}
            </div>
            <div className="w-full">
                <h5 className={`mb-3 text-lg font-semibold text-${type === "danger" ? "black dark:text-[#B45454]" : type}`}>
                    {title}
                </h5>
                <p className={`text-base leading-relaxed ${type === "danger" ? "text-[#CD5D5D]" : `text-${type}`}`}>
                    {message}
                </p>
            </div>
        </div>
    );
};

export default Alert;
