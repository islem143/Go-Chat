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
  <ScrollArea class="h-screen flex">
    <div v-if="users.length != 0" class="flex-1 flex flex-col gap-2 p-4 pt-0">
      <TransitionGroup name="list" appear>
        <button v-for="item of users" :key="item.id" :class="cn(
          'flex flex-col items-start gap-2 rounded-lg border p-3 text-left text-sm transition-all hover:bg-accent',
          selectedUser === item && 'bg-blue-200'
        )
          " @click="selectedUser = item">

          <div class="flex w-full flex-col gap-1">
     
            <div class="flex items-center">
              <div class="flex items-center gap-2">
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
               
              </div>
            </div>

           
          </div>
          <div class="line-clamp-2 text-xs text-muted-foreground">
       
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
  </ScrollArea>
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
