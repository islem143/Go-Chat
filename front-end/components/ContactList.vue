<script lang="ts" setup>
import type { User } from "~/types";
import { ScrollArea } from "@/components/ui/scroll-area";
import { cn } from "@/lib/utils";


interface UsersList {
  users: User[];
}

const props = defineProps<UsersList>();

const selectedUser = defineModel<string>("selectedUser", { required: true });



</script>

<template>
  <ScrollArea class="h-screen">
    <div v-if="users.length > 0" class="flex flex-col gap-2 p-4 pt-0">
      <TransitionGroup name="list" appear>
        <button 
          v-for="user in users" 
          :key="user.id" 
          :class="cn(
            'flex flex-col gap-2 w-full rounded-lg border p-3 text-left transition-all hover:bg-gray-50',
            selectedUser === user.id ? 'bg-blue-50 border-blue-200' : 'border-gray-200'
          )"
          @click="selectedUser = user"
        >
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <Avatar class="h-10 w-10">
                <AvatarImage src="/avatars/01.png" alt="User avatar" />
                <AvatarFallback>{{ user.name[0] }}</AvatarFallback>
              </Avatar>
              <div>
                <p class="font-medium text-sm">{{ user.name }}</p>
                <p class="text-xs text-gray-500">{{ user.email }}</p>
              </div>
            </div>
            <div class="text-xs text-gray-500">
              12m ago
            </div>
          </div>
          
          <div class="pl-[52px] text-sm text-gray-600 line-clamp-1">
            Last message preview goes here...
          </div>

          <div class="flex items-center gap-2 pl-[52px]">
            <span v-if="true" class="flex h-6 w-6 items-center justify-center rounded-full bg-blue-100 text-xs font-medium text-blue-600">
              3
            </span>
          </div>
        </button>
      </TransitionGroup>
    </div>
    
    <div v-else class="flex h-full items-center justify-center p-4">
      <div class="text-center">
        <p class="text-sm text-gray-500">No contacts found</p>
      </div>
    </div>
  </ScrollArea>
  <!-- <ScrollArea class="h-screen flex">
    <div v-if="users.length != 0" class="flex-1 flex flex-col gap-2 p-4 pt-0">
      <TransitionGroup name="list" appear>
        <button v-for="item of users" :key="item.id" :class="cn(
          'rounded-lg border p-3 text-left text-sm transition-all hover:bg-accent',
          selectedUser === item && 'bg-blue-100'
        )
          " @click="selectedUser = item">

          <div class="flex w-full flex-col gap-1">

            <div class="flex items-center">
              <div class="flex items-center gap-2">
                <Avatar>
                  <AvatarImage src="" alt="@radix-vue" />
                  <AvatarFallback>CN</AvatarFallback>
                </Avatar>
                <div class="font-semibold">
                  {{ item.name }}
                </div>

              </div>
              <div :class="cn(
                'ml-auto text-xs',
                selectedUser === item.id
                  ? 'text-foreground'
                  : 'text-muted-foreground'
              )
                ">
                last seen.
              </div>
            </div>


          </div>
          <div class="line-clamp-2 ml-10 text-xs text-muted-foreground">
            preview message.
          </div>
          <div class="flex items-center gap-2">

          </div>
        </button>
      </TransitionGroup>
    </div>
    <div v-else>
      <p class="p-4"> Users not found.
      </p>
    </div>
  </ScrollArea> -->
</template>

<style scoped>
.list-move,
.list-enter-active,
.list-leave-active {
  transition: all 0.5s ease;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateY(15px);
}

.list-leave-active {
  position: absolute;
}
</style>
