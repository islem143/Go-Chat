<script setup lang="ts">
import { Button } from "@/components/ui/button";
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

import { h } from "vue";
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";
import { Actions } from "~/constants/actions";

const formData = reactive({
  name: "",
  email: "",
  password: "",
});
const loading = useLoading();

const formSchema = toTypedSchema(
  z.object({
    name: z.string().min(2).max(50),
    email: z.string().nonempty("This is required").email(),
    password: z.string().nonempty("This is required").min(5),
  })
);
const { handleSubmit } = useForm({
  validationSchema: formSchema,
});

const signUp = handleSubmit((values) => {
  Auth.register({
    name: formData.name,
    email: formData.email,
    password: formData.password,
  }).then((res) => {
    navigateTo("/login");
  });
});
</script>

<template>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <Card class="w-full max-w-md">
      <CardHeader>
        <CardTitle class="text-2xl font-bold text-center text-gray-900">
          Create your account
        </CardTitle>
        <CardDescription class="text-center text-gray-600">
          Enter your information below to get started
        </CardDescription>
      </CardHeader>
      <CardContent>
       
          <div class="space-y-4">
            <FormField v-slot="{ componentField }" name="name">
              <FormItem>
                <FormLabel class="text-sm font-medium text-gray-700">Full Name</FormLabel>
                <FormControl>
                  <Input
                    v-model="formData.name"
                    type="text"
                    placeholder="John Doe"
                    class="w-full"
                    v-bind="componentField"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            </FormField>

            <FormField v-slot="{ componentField }" name="email">
              <FormItem>
                <FormLabel class="text-sm font-medium text-gray-700">Email address</FormLabel>
                <FormControl>
                  <Input
                    v-model="formData.email"
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
                    v-model="formData.password"
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
            :disabled="loading.loading[Actions.Register]"
            @click="signUp"
          >
            <Loader2 v-if="loading.loading[Actions.Register]" class="w-4 h-4 mr-2 animate-spin" />
            Create Account
          </Button>

          <p class="text-center text-sm text-gray-600">
            Already have an account?
            <NuxtLink to="/login" class="font-medium text-blue-600 hover:text-blue-500">
              Sign in
            </NuxtLink>
          </p>
 
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
        <div class="grid grid-cols-2 gap-4">
          <div class="grid gap-2">
            <FormField v-slot="{ componentField }" name="name">
              <FormItem>
                <FormLabel>name</FormLabel>
                <FormControl>
                  <Input
                    v-model="formData.name"
                    type="text"
                    placeholder="name"
                    v-bind="componentField"
                  />
                </FormControl>

                <FormMessage />
              </FormItem>
            </FormField>
          </div>
        </div>
        <div class="grid gap-2">
          <FormField v-slot="{ componentField }" name="email">
            <FormItem>
              <FormLabel>email</FormLabel>
              <FormControl>
                <Input
                  v-model="formData.email"
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
                <Input
                  v-model="formData.password"
                  type="password"
                  placeholder="password"
                  v-bind="componentField"
                />
              </FormControl>

              <FormMessage />
            </FormItem>
          </FormField>
        </div>
        <Button :disabled="loading.loading[Actions.Register]" @click="signUp" class="w-full">
           Create an account
           <Loader2 v-if="loading.loading[Actions.Register]" class="w-4 h-4 mr-2 animate-spin" />

          </Button>
        <Button variant="outline" class="w-full"> Sign up with GitHub </Button>
      </div>
      <div class="mt-4 text-center text-sm">
        Already have an account?

        <NuxtLink to="/login" class="underline"> Sign in</NuxtLink>
      </div>
    </CardContent>
  </Card> -->
</template>
