import { configureStore } from "@reduxjs/toolkit";

import { mediaApi } from "@/redux/services/mediaApi";
import { setupListeners } from "@reduxjs/toolkit/dist/query";
import { recipientApi } from "@/redux/services/recipientsApi";
import { newslettersApi } from "@/redux/services/newslettersApi";

export const store = configureStore({
    reducer: {
        [mediaApi.reducerPath]: mediaApi.reducer,
        [recipientApi.reducerPath]: recipientApi.reducer,
        [newslettersApi.reducerPath]: newslettersApi.reducer,
    },
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware().concat([
            mediaApi.middleware,
            recipientApi.middleware,
            newslettersApi.middleware,
        ])
});

setupListeners(store.dispatch);

// types autocomplete
export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch