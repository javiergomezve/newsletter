"use client";

import { useCallback } from "react";

interface ModalProps {
    isOpen?: boolean;
    title?: string;
    body?: React.ReactElement;
    footer?: React.ReactElement;
    disabled?: boolean;
    onClose: () => void;
    onSubmit: () => void;
}

const Modal: React.FC<ModalProps> = (props) => {
    const {
        isOpen, title, body, footer, disabled, onClose, onSubmit
    } = props;

    const handleClose = useCallback(() => {
        if (disabled) return;

        onClose();
    }, [disabled, onClose]);

    const handleSubmit = useCallback(() => {
        if (disabled) return;

        onSubmit();
    }, [disabled, onSubmit]);

    if (!isOpen) return null;

    return (
        <div
            className="fixed top-0 left-0 z-999999 flex h-full min-h-screen w-full items-center justify-center bg-black/90 px-4 py-5">
            <div
                className="w-full max-w-142.5 rounded-lg bg-white py-12 px-8 text-center
            dark:bg-boxdark md:py-15 md:px-17.5">
                <h3 className="pb-2 text-xl font-bold text-black dark:text-white sm:text-2xl">
                    {title}
                </h3>
                <span className="mx-auto mb-6 inline-block h-1 w-22.5 rounded bg-primary"></span>

                {body}

                <br/>

                <div className="-mx-3 flex flex-wrap gap-y-4">
                    <div className="w-full px-3 2xsm:w-1/2">
                        <button
                            className="block w-full rounded border border-stroke bg-gray p-3 text-center font-medium text-black transition hover:border-meta-1 hover:bg-meta-1 hover:text-white dark:border-strokedark dark:bg-meta-4 dark:text-white dark:hover:border-meta-1 dark:hover:bg-meta-1"
                            onClick={handleClose}
                        >
                            Cancel
                        </button>
                    </div>
                    <div className="w-full px-3 2xsm:w-1/2">
                        <button
                            className="block w-full rounded border border-primary bg-primary p-3 text-center font-medium text-white transition hover:bg-opacity-90"
                            onClick={handleSubmit}
                        >
                            Confirm
                        </button>
                    </div>
                </div>
            </div>
        </div>
    )
        ;
};

export default Modal;