<script lang="ts" setup>
import {
  Archive,
  ArchiveX,
  Clock,
  Forward,
  MoreVertical,
  Reply,
  ReplyAll,
  Trash2,
} from "lucide-vue-next";
import { computed } from "vue";
import addDays from "date-fns/addDays";
import addHours from "date-fns/addHours";
import format from "date-fns/format";
import nextSaturday from "date-fns/nextSaturday";
import type { Mail } from "../data/mails";
import { Calendar } from "@/components/ui/calendar";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { Separator } from "@/components/ui/separator";
import { Switch } from "@/components/ui/switch";
import { Textarea } from "@/components/ui/textarea";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";

interface MailDisplayProps {
  mail: Mail | undefined;
}

const mailFallbackName = computed(() => {
  return props.mail?.name
    .split(" ")
    .map((chunk) => chunk[0])
    .join("");
});

const today = new Date();

import { Check, Plus, Send } from "lucide-vue-next";
import { computed, ref } from "vue";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from "@/components/ui/card";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";
import { Message } from "~/api/Message";

const input = ref("");
const inputLength = computed(() => input.value.trim().length);

let socket = new WebSocket("ws://localhost:8000/ws/private-chat/4");
socket.addEventListener("open", (event) => {
  console.log("connection oppend");
  let msg = {
    user_id:
      props.selectedUser.name == "ilyes"
        ? "664f842a19d94fd5534d2eac"
        : "664f842019d94fd5534d2eab",
  };
  socket.send(JSON.stringify(msg));
});
const props = defineProps(["selectedUser"]);

const sendMessage = () => {
  messages.value.push({
    role: "user",
    content: input.value,
  });
  console.log(props.selectedUser.id, "sdddddd");

  let msg = {
    user_id:
      props.selectedUser.name == "ilyes"
        ? "664f842019d94fd5534d2eab"
        : "664f842a19d94fd5534d2eac",
    message: input.value,
    receiver: props.selectedUser.id,
  };
  socket.send(JSON.stringify(msg));
  input.value = "";
};

socket.onmessage = (msg) => {
  console.log("get Message***********");
  messages.value.push({ role: "is", content: JSON.parse(msg.data) });
  //console.log(r);

  //messages.value.push({role:"ss",content})
};


let users = ref<any>([]);
let messages = ref<any>([]);

type User = (typeof users.value)[number];
watch(props,async (val)=>{

  let res=await Message.getChatMessages({receiverId:props.selectedUser.id});
  res=res.map(p=>({"content":p.text,role:"qsd"}));

  
 messages.value=res;
  
})
const open = ref(false);
const selectedUsers = ref<User[]>([]);
</script>

<template>
  <div class="flex h-full flex-col">
    <div class="flex items-center p-2">
      <div class="flex items-center gap-2">
        <Tooltip>
          <TooltipTrigger as-child>
            <Button variant="ghost" size="icon" :disabled="!mail">
              <Archive class="size-4" />
              <span class="sr-only">Archive</span>
            </Button>
          </TooltipTrigger>
          <TooltipContent>Archive</TooltipContent>
        </Tooltip>
        <Tooltip>
          <TooltipTrigger as-child>
            <Button variant="ghost" size="icon" :disabled="!mail">
              <ArchiveX class="size-4" />
              <span class="sr-only">Move to junk</span>
            </Button>
          </TooltipTrigger>
          <TooltipContent>Move to junk</TooltipContent>
        </Tooltip>
        <Tooltip>
          <TooltipTrigger as-child>
            <Button variant="ghost" size="icon" :disabled="!mail">
              <Trash2 class="size-4" />
              <span class="sr-only">Move to trash</span>
            </Button>
          </TooltipTrigger>
          <TooltipContent>Move to trash</TooltipContent>
        </Tooltip>
        <Separator orientation="vertical" class="mx-1 h-6" />
        <Tooltip>
          <Popover>
            <PopoverTrigger as-child>
              <TooltipTrigger as-child>
                <Button variant="ghost" size="icon" :disabled="!mail">
                  <Clock class="size-4" />
                  <span class="sr-only">Snooze</span>
                </Button>
              </TooltipTrigger>
            </PopoverTrigger>
            <PopoverContent class="flex w-[535px] p-0">
              <div class="flex flex-col gap-2 border-r px-2 py-4">
                <div class="px-4 text-sm font-medium">Snooze until</div>
                <div class="grid min-w-[250px] gap-1">
                  <Button variant="ghost" class="justify-start font-normal">
                    Later today
                    <span class="ml-auto text-muted-foreground">
                      {{ format(addHours(today, 4), "E, h:m b") }}
                    </span>
                  </Button>
                  <Button variant="ghost" class="justify-start font-normal">
                    Tomorrow
                    <span class="ml-auto text-muted-foreground">
                      {{ format(addDays(today, 1), "E, h:m b") }}
                    </span>
                  </Button>
                  <Button variant="ghost" class="justify-start font-normal">
                    This weekend
                    <span class="ml-auto text-muted-foreground">
                      {{ format(nextSaturday(today), "E, h:m b") }}
                    </span>
                  </Button>
                  <Button variant="ghost" class="justify-start font-normal">
                    Next week
                    <span class="ml-auto text-muted-foreground">
                      {{ format(addDays(today, 7), "E, h:m b") }}
                    </span>
                  </Button>
                </div>
              </div>
              <div class="p-2">
                <Calendar />
              </div>
            </PopoverContent>
          </Popover>
          <TooltipContent>Snooze</TooltipContent>
        </Tooltip>
      </div>
      <div class="ml-auto flex items-center gap-2">
        <Tooltip>
          <TooltipTrigger as-child>
            <Button variant="ghost" size="icon" :disabled="!mail">
              <Reply class="size-4" />
              <span class="sr-only">Reply</span>
            </Button>
          </TooltipTrigger>
          <TooltipContent>Reply</TooltipContent>
        </Tooltip>
        <Tooltip>
          <TooltipTrigger as-child>
            <Button variant="ghost" size="icon" :disabled="!mail">
              <ReplyAll class="size-4" />
              <span class="sr-only">Reply all</span>
            </Button>
          </TooltipTrigger>
          <TooltipContent>Reply all</TooltipContent>
        </Tooltip>
        <Tooltip>
          <TooltipTrigger as-child>
            <Button variant="ghost" size="icon" :disabled="!mail">
              <Forward class="size-4" />
              <span class="sr-only">Forward</span>
            </Button>
          </TooltipTrigger>
          <TooltipContent>Forward</TooltipContent>
        </Tooltip>
      </div>
      <Separator orientation="vertical" class="mx-2 h-6" />
      <DropdownMenu>
        <DropdownMenuTrigger as-child>
          <Button variant="ghost" size="icon" :disabled="!mail">
            <MoreVertical class="size-4" />
            <span class="sr-only">More</span>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuItem>Mark as unread</DropdownMenuItem>
          <DropdownMenuItem>Star thread</DropdownMenuItem>
          <DropdownMenuItem>Add label</DropdownMenuItem>
          <DropdownMenuItem>Mute thread</DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
    <Separator />
    <Card class="overflow-scroll overflow-x-hidden">
      <CardHeader class="flex flex-row items-center justify-between">
        <div class="flex items-center space-x-4">
          <Avatar>
            <AvatarImage src="/avatars/01.png" alt="Image" />
            <AvatarFallback>OM</AvatarFallback>
          </Avatar>
          <div>
            <p class="text-sm font-medium leading-none">Sofia Davis</p>
            <p class="text-sm text-muted-foreground">m@example.com</p>
          </div>
        </div>
        <TooltipProvider>
          <Tooltip :delay-duration="200">
            <TooltipTrigger as-child>
              <Button
                variant="outline"
                class="rounded-full p-2.5 flex items-center justify-center"
                @click="open = true"
              >
                <Plus class="w-4 h-4" />
              </Button>
            </TooltipTrigger>
            <TooltipContent :side-offset="10"> New message </TooltipContent>
          </Tooltip>
        </TooltipProvider>
      </CardHeader>
      <CardContent>
        <div class="space-y-4 ">

          <div
            v-for="(message, index) in messages"
            :key="index"
            :class="
              cn(
                'flex w-max max-w-[75%] flex-col gap-2 rounded-lg px-3 py-2 text-sm',
                message.role === 'user'
                  ? 'ml-auto bg-primary text-primary-foreground'
                  : 'bg-muted'
              )
            "
          >
            {{ message.content }}
          </div>
        </div>
      </CardContent>
      <CardFooter>
        <form
          class="flex w-full items-center space-x-2"
          @submit.prevent="sendMessage"
        >
          <Input
            v-model="input"
            placeholder="Type a message..."
            class="flex-1"
          />
          <Button
            class="p-2.5 flex items-center justify-center"
            :disabled="inputLength === 0"
          >
            <Send class="w-4 h-4" />
            <span class="sr-only">Send</span>
          </Button>
        </form>
      </CardFooter>
    </Card>

    <Dialog v-model:open="open">
      <DialogContent class="gap-0 p-0 outline-none">
        <DialogHeader class="px-4 pb-4 pt-5">
          <DialogTitle>New message</DialogTitle>
          <DialogDescription>
            Invite a user to this thread. This will create a new group message.
          </DialogDescription>
        </DialogHeader>
        <Command
          class="overflow-hidden rounded-t-none border-t"
          :filter-function="(list: User[], search) => list.filter(l => l.name.toLowerCase().includes(search.toLowerCase()))"
        >
          <CommandInput placeholder="Search user..." />
          <CommandList>
            <CommandEmpty>No users found.</CommandEmpty>
            <CommandGroup class="p-2">
              <CommandItem
                v-for="user in users"
                :key="user.email"
                :value="user"
                class="flex items-center px-2"
                @select="
                  () => {
                    const index = selectedUsers.findIndex((u) => u === user);
                    if (index !== -1) {
                      selectedUsers.splice(index, 1);
                    } else {
                      selectedUsers.push(user);
                    }
                  }
                "
              >
                <Avatar>
                  <AvatarImage :src="user.avatar" alt="Image" />
                  <AvatarFallback>{{ user.name[0] }}</AvatarFallback>
                </Avatar>
                <div class="ml-2">
                  <p class="text-sm font-medium leading-none">
                    {{ user.name }}
                  </p>
                  <p class="text-sm text-muted-foreground">
                    {{ user.email }}
                  </p>
                </div>
                <Check
                  v-if="selectedUsers.includes(user)"
                  class="ml-auto flex h-5 w-5 text-primary"
                />
              </CommandItem>
            </CommandGroup>
          </CommandList>
        </Command>
        <DialogFooter class="flex items-center border-t p-4 sm:justify-between">
          <div
            v-if="selectedUsers.length > 0"
            class="flex -space-x-2 overflow-hidden"
          >
            <Avatar
              v-for="user in selectedUsers"
              :key="user.email"
              class="inline-block border-2 border-background"
            >
              <AvatarImage :src="user.avatar" />
              <AvatarFallback>{{ user.name[0] }}</AvatarFallback>
            </Avatar>
          </div>

          <p v-else class="text-sm text-muted-foreground">
            Select users to add to this thread.
          </p>

          <Button :disabled="selectedUsers.length < 2" @click="open = false">
            Continue
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  </div>
</template>
../components/data/mails
