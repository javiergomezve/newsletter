import React from "react";

interface CheckboxProps {
    label: string;
    isChecked: boolean;
    onChange: (checked: boolean) => void;
    containerClassName?: string;
}

const Checkbox: React.FC<CheckboxProps> = ({ label, isChecked, onChange, containerClassName }) => {
    const handleCheckboxChange = () => {
        onChange(!isChecked);
    };

    return (
        <div className={containerClassName}>
            <label
                htmlFor={label}
                className="flex cursor-pointer select-none items-center"
            >
                <div className="relative">
                    <input
                        type="checkbox"
                        id={label}
                        className="sr-only"
                        onChange={handleCheckboxChange}
                        checked={isChecked}
                    />
                    <div
                        className={`mr-4 flex h-5 w-5 items-center justify-center rounded border ${
                            isChecked && "border-primary bg-gray dark:bg-transparent"
                        }`}
                    >
                        <span className={`h-2.5 w-2.5 rounded-sm ${isChecked && "bg-primary"}`}></span>
                    </div>
                </div>

                {label}
            </label>
        </div>
    );
};

export default Checkbox;
