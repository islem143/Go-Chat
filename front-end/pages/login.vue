<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { Loader2 } from "lucide-vue-next";
import { Actions } from "~/constants/actions";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

import { h,ref } from "vue";
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";
import { useToast } from "@/components/ui/toast/use-toast";
import { Toaster } from "@/components/ui/toast";
import { useAuth } from "~/store/auth";

const auth = useAuth();
const email = ref('')
const password = ref('')
watch(email,(val)=>{
  console.log(val,"----");
  
})
// const formSchema = toTypedSchema(
//   z.object({
//     email: z.string().nonempty("This is required").email(),
//     password: z.string().nonempty("This is required").min(5),
//   })
// );
// const { handleSubmit } = useForm({
//   validationSchema: formSchema,
// });
const loading = useLoading();


const login = ()=>{
  console.log("email");
  
  auth
    .login({
      email: email.value,
      password: password.value,
    })
    .then(() => {
      navigateTo("/");
    });
}





</script>

<template>
  <Card class="mx-auto mt-8">
    <CardHeader>
      <CardTitle class="text-xl"> Sign Up </CardTitle>
      <CardDescription>
        Enter your information to create an account
      </CardDescription>
    </CardHeader>
    <CardContent>
      <div class="grid gap-4">
        <div class="grid gap-2">
          <FormField v-slot="{ componentField }" name="email">
            <FormItem>
              <FormLabel>email</FormLabel>
              <FormControl>
                <input
                  v-model="email"
                  type="text"
                  placeholder="email"
                  v-bind="componentField"
                />
              </FormControl>

              <FormMessage />
            </FormItem>
          </FormField>
        </div>
        <div class="grid gap-2">
          <FormField v-slot="{ componentField }" name="password">
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <input
                  v-model="password"
                  type="password"
                  placeholder="password"
                  v-bind="componentField"
                />
              </FormControl>

              <FormMessage />
            </FormItem>
          </FormField>
        </div>
        <Button
          :disabled="loading.loading[Actions.Login]"
          @click="login"
          class="w-full"
        >
          <Loader2
            v-if="loading.loading[Actions.Login]"
            class="w-4 h-4 mr-2 animate-spin"
          />
          Login
        </Button>
      </div>
      <div class="mt-4 text-center text-sm">
        D'ont have an account?

        <NuxtLink to="/register" class="underline"> Sign up</NuxtLink>
      </div>
    </CardContent>
  </Card>
</template>
