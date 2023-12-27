import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export interface Newsletter {
    id?: string;
    subject: string;
    content: string;
    send_at: string;
    recipients?: { [recipientId: string]: boolean };
    attachments?: { [attachmentId: string]: boolean };
}

export const newslettersApi = createApi({
    reducerPath: "newsletterAPI",
    baseQuery: fetchBaseQuery({
        baseUrl: process.env.NEXT_PUBLIC_API_URL
    }),
    endpoints: (builder) => ({
        getNewsLetters: builder.query<Newsletter[], null>({
            query: () => "newsletters",
            transformResponse: (response: { error: boolean; data: Newsletter[] }) => response.data
        }),
        createNewsletter: builder.mutation<Newsletter, Newsletter>({
            query: (newsletter) => {
                let recipients: string[] = [];
                if (newsletter.recipients) {
                    recipients = Object.keys(newsletter.recipients);
                }

                let attachments: string[] = [];
                if (newsletter.attachments) {
                    attachments = Object.keys(newsletter.attachments);
                }

                return {
                    url: "newsletters",
                    method: "POST",
                    body: {
                        ...newsletter,
                        send_at: new Date(),
                        recipients,
                        attachments
                    },
                    headers: {
                        "Content-Type": "application/json"
                    }
                }
            },
            transformResponse: (response: { error: boolean; data: Newsletter }) => response.data
        })
    })
});

export const {
    useGetNewsLettersQuery,
    useCreateNewsletterMutation
} = newslettersApi;