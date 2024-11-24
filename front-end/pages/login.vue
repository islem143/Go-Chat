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

import { h, ref } from "vue";
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";
import { useToast } from "@/components/ui/toast/use-toast";
import { Toaster } from "@/components/ui/toast";
import { useAuth } from "~/store/auth";

const auth = useAuth();
const email = ref('')
const password = ref('')
watch(email, (val) => {
  console.log(val, "----");

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


const login = () => {
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
  <div class="min-h-screen bg-gray-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <Card class="w-full max-w-md">
      <CardHeader>
        <CardTitle class="text-2xl font-bold text-center text-gray-900">
          Welcome back
        </CardTitle>
        <CardDescription class="text-center text-gray-600">
          Sign in to your account to continue
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form @submit.prevent="login" class="space-y-6">
          <div class="space-y-4">
            <FormField v-slot="{ componentField }" name="email">
              <FormItem>
                <FormLabel class="text-sm font-medium text-gray-700">Email address</FormLabel>
                <FormControl>
                  <Input
                    v-model="email"
                    type="email"
                    placeholder="you@example.com"
                    class="w-full"
                    v-bind="componentField"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            </FormField>

            <FormField v-slot="{ componentField }" name="password">
              <FormItem>
                <FormLabel class="text-sm font-medium text-gray-700">Password</FormLabel>
                <FormControl>
                  <Input
                    v-model="password"
                    type="password"
                    placeholder="••••••••"
                    class="w-full"
                    v-bind="componentField"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            </FormField>
          </div>

          <Button 
            type="submit"
            class="w-full bg-blue-600 hover:bg-blue-700 text-white"
            :disabled="loading.loading[Actions.Login]"
          >
            <Loader2 v-if="loading.loading[Actions.Login]" class="w-4 h-4 mr-2 animate-spin" />
            Sign in
          </Button>

          <p class="text-center text-sm text-gray-600">
            Don't have an account?
            <NuxtLink to="/register" class="font-medium text-blue-600 hover:text-blue-500">
              Sign up
            </NuxtLink>
          </p>
        </form>
      </CardContent>
    </Card>
  </div>
  <!-- <Card class="mx-auto mt-8">
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
                <input v-model="email" type="text" placeholder="email" v-bind="componentField" />
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
                <input v-model="password" type="password" placeholder="password" v-bind="componentField" />
              </FormControl>

              <FormMessage />
            </FormItem>
          </FormField>
        </div>
        <Button :disabled="loading.loading[Actions.Login]" @click="login" class="w-full">
          <Loader2 v-if="loading.loading[Actions.Login]" class="w-4 h-4 mr-2 animate-spin" />
          Login
        </Button>
      </div>
      <div class="mt-4 text-center text-sm">
        D'ont have an account?

        <NuxtLink to="/register" class="underline"> Sign up</NuxtLink>
      </div>
    </CardContent>
  </Card> -->
</template>
