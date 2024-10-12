<script lang="ts" setup>
import { formatDistanceToNow } from "date-fns";
import type { User } from "~/types";
import { Message } from "~/api/Message";
import { ScrollArea } from "@/components/ui/scroll-area";
import { cn } from "@/lib/utils";
import { Badge } from "@/components/ui/badge";

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
            <!-- <Avatar>
          <AvatarImage src="/avatars/01.png" alt="Image" />
          <AvatarFallback>OM</AvatarFallback>
        </Avatar> -->
            <div class="flex items-center">
              <div class="flex items-center gap-2">
                <div class="font-semibold">
                  {{ item.name }}
                </div>
                <!-- <span
                  v-if="!item.read"
                  class="flex h-2 w-2 rounded-full bg-blue-600"
                /> -->
              </div>
              <div :class="cn(
                'ml-auto text-xs',
                selectedUser === item.id
                  ? 'text-foreground'
                  : 'text-muted-foreground'
              )
                ">
                <!-- {{
                  formatDistanceToNow(new Date(item.date), { addSuffix: true })
                }} -->
              </div>
            </div>

            <!-- <div class="text-xs font-medium">
              {{ item.subject }}
            </div> -->
          </div>
          <div class="line-clamp-2 text-xs text-muted-foreground">
            <!-- {{ item.text.substring(0, 50) }} -->
          </div>
          <div class="flex items-center gap-2">
            <!-- <Badge v-for="label of item.labels" :key="label" :variant="getBadgeVariantFromLabel(label)">
              {{ label }}
            </Badge> -->
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
