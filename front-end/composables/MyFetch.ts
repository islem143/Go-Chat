import { useToast } from "@/components/ui/toast/use-toast";
import type { Actions } from "~/constants/actions";
import { useLoading } from "~/store/loading";

export const useMyFetch = async (url:string, opts:object, type:Actions) => {
  const config = useRuntimeConfig();
  const { toast } = useToast();
  const loading = useLoading();
  loading.loading[type] = true;
  const { pending, data, error, status } = await useFetch(url, {
    baseURL: config.public.baseURL,
    credentials: 'include',
    ...opts,
  });
  loading.loading[type] = false;

  if (error.value) {
    let message;

    switch (error.value.statusCode) {
      case 404:
      case 422:
        message = error.value.data.message;
        break;
      case 401:
        message = "Unauthorized"
        break;
      case 403:
        message = "Forbiden"
        break;
      case 500:
        message = "server error";
        break;
    }
    toast({
      title: message,

      variant: "destructive",
    });
    return { data: null, error: error.value };
  }

  return { data: data.value, error: null };
};
