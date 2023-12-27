import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
export interface ErrorResponse {
    data: {
        message: string
    }
}

export interface Media {
    id: string;
    file_name: string;
    content_type: string;
    location: string;
}

export const mediaApi = createApi({
    reducerPath: "mediaAPI",
    baseQuery: fetchBaseQuery({
        baseUrl: process.env.NEXT_PUBLIC_API_URL
    }),
    tagTypes: ["Medias"],
    endpoints: (builder) => ({
        getMedias: builder.query<Media[], null>({
            query: () => "/media",
            transformResponse: (response: { error: boolean; data: Media[] }) => response.data,
        }),
        getMediaById: builder.query<Media, { id: string }>({
            query: ({ id }) => `/media/${id}`,
            transformResponse: (response: { error: boolean; data: Media }) => response.data
        }),
        createMedia: builder.mutation<Media, FormData>({
            query: (formData) => ({
                url: "/media",
                method: "POST",
                body: formData
            }),
            transformResponse: (response: { error: boolean; data: Media }) =>
                response.data
        }),
        updateMedia: builder.mutation<Media, { id: string; formData: FormData }>({
            query: ({ id, formData }) => ({
                url: `/media/${id}`,
                method: "PUT",
                body: formData
            }),
            transformResponse: (response: { error: boolean; data: Media }) =>
                response.data
        }),
        deleteMedia: builder.mutation<void, { id: string }>({
            query: ({ id }) => ({
                url: `/media/${id}`,
                method: "DELETE"
            })
        })
    })
});

export const {
    useGetMediasQuery,
    useGetMediaByIdQuery,
    useCreateMediaMutation,
    useUpdateMediaMutation,
    useDeleteMediaMutation
} = mediaApi;