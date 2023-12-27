import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export interface Recipient {
    id?: string;
    full_name: string;
    email: string;
    status?: string;
}

export const recipientApi = createApi({
    reducerPath: "recipientAPI",
    baseQuery: fetchBaseQuery({
        baseUrl: process.env.NEXT_PUBLIC_API_URL
    }),
    endpoints: (builder) => ({
        getRecipients: builder.query<Recipient[], null>({
            query: () => "/recipients",
            transformResponse: (response: { error: boolean; data: Recipient[] }) => response.data,
        }),
        createRecipients: builder.mutation<Recipient[], Recipient[]>({
            query: (recipients) => ({
                url: "recipients",
                method: "POST",
                body: recipients,
                headers: {
                    'Content-Type': 'application/json',
                },
            }),
            transformResponse: (response: { error: boolean; data: Recipient[] }) => response.data,
        }),
    }),
});

export const {
    useGetRecipientsQuery,
    useCreateRecipientsMutation,
} = recipientApi;