<script setup>
import { Button } from '@/components/ui/button'
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <div class="container mx-auto px-4 py-16">
      <div class="flex flex-col items-center text-center">
        <h1 class="text-4xl md:text-6xl font-bold text-gray-900 mb-6">
          Connect Instantly with Real-Time Chat
        </h1>
        <p class="text-xl text-gray-600 mb-8 max-w-2xl">
          Experience seamless communication with our modern chat application. Send messages, share moments, and stay connected with friends and colleagues.
        </p>
        <div class="flex gap-4">
          <NuxtLink to="/chat">
            <Button class="bg-blue-500 hover:bg-blue-600 text-white px-8 py-2 rounded-full text-lg">
              Start Chatting
            </Button>
          </NuxtLink>
          <NuxtLink to="/login">
            <Button variant="outline" class="px-8 py-2 rounded-full text-lg">
              Sign In
            </Button>
          </NuxtLink>
        </div>
      </div>

      <div class="mt-20 grid md:grid-cols-3 gap-8">
        <div class="bg-white p-6 rounded-lg shadow-sm">
          <div class="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mb-4">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-blue-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
            </svg>
          </div>
          <h3 class="text-xl font-semibold mb-2">Real-Time Messaging</h3>
          <p class="text-gray-600">Instant message delivery with real-time typing indicators and read receipts.</p>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-sm">
          <div class="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mb-4">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-blue-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
          </div>
          <h3 class="text-xl font-semibold mb-2">Secure Communication</h3>
          <p class="text-gray-600">End-to-end encryption ensures your conversations remain private and secure.</p>
        </div>

        <div class="bg-white p-6 rounded-lg shadow-sm">
          <div class="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mb-4">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-blue-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
          </div>
          <h3 class="text-xl font-semibold mb-2">Group Chats</h3>
          <p class="text-gray-600">Create and manage group conversations with multiple participants easily.</p>
        </div>
      </div>
    </div>
  </div>
</template>
