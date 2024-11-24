<script setup>
import { Button } from '@/components/ui/button'
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from '@/components/ui/dropdown-menu'
import { Icon } from '@iconify/vue'

const colorMode = useColorMode();
colorMode.preference='light'
</script>
<template>
  <div class="flex min-h-screen w-full flex-col">
    <header class="sticky top-0 z-50 flex h-16 items-center border-b bg-white shadow-sm">
      <div class="container mx-auto flex items-center justify-between px-4">
        <div class="flex items-center gap-6">
          <NuxtLink to="/" class="flex items-center gap-2 text-xl font-bold text-blue-600">
            <Icon icon="heroicons:chat-bubble-left-right" class="h-6 w-6" />
            ChatApp
          </NuxtLink>
          
          <nav class="hidden md:flex items-center gap-6">
            <NuxtLink 
              to="/" 
              class="text-sm font-medium text-gray-600 hover:text-blue-600 transition-colors"
            >
              Home
            </NuxtLink>
            <NuxtLink 
              to="/chat" 
              class="text-sm font-medium text-gray-600 hover:text-blue-600 transition-colors"
            >
              Chat
            </NuxtLink>
          </nav>
        </div>

        <div class="flex items-center gap-4">
          <NuxtLink 
            to="/login"
            class="hidden md:inline-flex text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 px-4 py-2 rounded-full transition-colors"
          >
            Sign In
          </NuxtLink>

          <!-- Mobile Menu -->
          <Button variant="ghost" size="icon" class="md:hidden" @click="() => {}">
            <Icon icon="heroicons:bars-3" class="h-6 w-6" />
          </Button>
        </div>
      </div>
    </header>

    <main class="flex-1">
      <slot />
    </main>
  </div>
  <!-- <div class="flex min-h-screen w-full flex-col">
    <header class="sticky top-0 flex h-16 items-center gap-4 border-b bg-background px-4 md:px-6">
      <nav
        class="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6">
        <a class="flex items-center gap-2 text-lg font-semibold md:text-base">
       
        </a>
      <a href="#" class="text-foreground transition-colors hover:text-foreground">
        <NuxtLink to="/">Home</NuxtLink>
      </a>

        <a href="#" class="text-foreground transition-colors hover:text-foreground">
          <NuxtLink to="/chat">Chat</NuxtLink>
        </a>
        <a href="#" class="text-foreground transition-colors hover:text-foreground">
          <NuxtLink to="/login">Login</NuxtLink>
        </a>

    
      </nav>
      <Sheet>
        <SheetTrigger as-child>
          <Button variant="outline" size="icon" class="shrink-0 md:hidden">
            <Menu class="h-5 w-5" />
            <span class="sr-only">Toggle navigation menu</span>
          </Button>
        </SheetTrigger>
        <SheetContent side="left">
          <nav class="grid gap-6 text-lg font-medium">
            <a href="#" class="flex items-center gap-2 text-lg font-semibold">
              <Package2 class="h-6 w-6" />
              <span class="sr-only">Acme Inc</span>
            </a>
            <a href="#" class="hover:text-foreground">
              Dashboard
            </a>
            <a href="#" class="text-muted-foreground hover:text-foreground">
              Orders
            </a>
            <a href="#" class="text-muted-foreground hover:text-foreground">
              Products
            </a>
            <a href="#" class="text-muted-foreground hover:text-foreground">
              Customers
            </a>
            <a href="#" class="text-muted-foreground hover:text-foreground">
              Analytics
            </a>
          </nav>
        </SheetContent>
      </Sheet>
      <div class="flex w-full items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
        <form class="ml-auto flex-1 sm:flex-initial">

        </form>
        <DropdownMenu>
          <DropdownMenuTrigger as-child>
            <Button variant="secondary" size="icon" class="rounded-full">
              <CircleUser class="h-5 w-5" />
              <span class="sr-only">Toggle user menu</span>
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuLabel>My Account</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem>Settings</DropdownMenuItem>
            <DropdownMenuItem>Support</DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem>Logout</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </header>

    <slot />
  </div> -->
</template>
