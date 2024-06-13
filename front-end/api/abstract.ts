import type { Actions } from "~/constants/actions";

export default class Abstract {
    static async get(url: string, type: Actions, query: any) {
        try {
            const config = useRuntimeConfig();

            const { data, error } = await useMyFetch(url, {
                baseURL: config.public.baseURL,

                method: "GET",
                query: query


            }, type)


            if (error) {
                throw error;
            }
            return data




        } catch (err) {
            throw err;
        }
    }
    static async post(url: string, body: any, type: Actions) {

        try {
            const config = useRuntimeConfig();

            const { data, error } = await useMyFetch(url, {
                baseURL: config.public.baseURL,

                method: "POST",
                body: body

            }, type)


            if (error) {
                throw error;
            }
            return data




        } catch (err) {
            throw err;
        }

    };


}

