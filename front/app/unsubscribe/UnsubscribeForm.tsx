"use client";

import { FormEvent, useCallback } from "react";

import { useUnsubscribeRecipientMutation } from "@/redux/services/recipientsApi";
import Spinner from "@/components/Spinner";

interface Props {
    email?: string;
}

function sleep(seconds: number) {
    return new Promise(resolve => setTimeout(resolve, seconds * 1000));
}

const UnsubscribeForm = ({ email }: Props) => {
    const [
        unsubscribe,
        { isLoading, isError, isSuccess }
    ] = useUnsubscribeRecipientMutation();

    const handleSubmit = useCallback(async (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        if (!email) return;

        await unsubscribe(email);

    }, [email]);


    return (
        <form onSubmit={handleSubmit}>
            <div className="p-6.5">
                <p className="mb-2">
                    We are sorry to see you go <strong>{email}</strong>
                </p>

                {isError && (
                    <p className="mb-2">
                        Validation errors
                    </p>
                )}

                {!isSuccess && (
                    <button
                        className="flex w-full justify-center items-center rounded bg-primary p-3 font-medium text-gray"
                        disabled={isLoading}
                    >
                        {isLoading && <Spinner />}

                        Confirm
                    </button>
                )}

                {isSuccess && (
                    <p>Subscription ended</p>
                )}
            </div>
        </form>
    );
};

export default UnsubscribeForm;