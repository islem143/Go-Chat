<script lang="ts" setup>
import { Send } from "lucide-vue-next";

import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from "@/components/ui/card";
import EmojiPicker from 'vue3-emoji-picker'
import 'vue3-emoji-picker/css'

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";

import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";
import { Message } from "~/api/Message";

import { computed } from "vue";


import { Button } from "@/components/ui/button";


function onSelectEmoji(emoji) {
  console.log(emoji)
  /*
    // result
    { 
        i: "😚", 
        n: ["kissing face"], 
        r: "1f61a", // with skin tone
        t: "neutral", // skin tone
        u: "1f61a" // without tone
    }
    */
}
const input = ref("");
const showTyping = ref(false);
let timeoutId = null;



const inputLength = computed(() => input.value.trim().length);

let socket = new WebSocket("ws://localhost:8000/ws/private-chat/4");
socket.addEventListener("open", (event) => {

});
const props = defineProps(["selectedUser"]);

const sendMessage = (event, type = 'message') => {



  let msg = {

    message: type == 'message' ? input.value : '',
    receiver: props.selectedUser.id,
    type: type,

  };



  socket.send(JSON.stringify(msg));
  if (type != "typing") {
    messages.value.push({
      role: "user",
      content: input.value,

    });
    input.value = "";
  }



};

socket.onmessage = (msg) => {
  let res = JSON.parse(msg.data);
  if (res.type == "typing") {
    showTyping.value = true;

    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => {

      showTyping.value = false;
    }, 1000);
  } else {

    messages.value.push({ role: "other", content: res.text, type: 'message' });
  }


};

const getMessages = async () => {
  let res = await Message.getChatMessages({ receiverId: props.selectedUser.id });


  res = res.map(p => ({ "content": p.text, role: "qsd", role: props.selectedUser.id == p.receiver_id ? 'user' : 'other' }));


  messages.value = res;
};

let messages = ref<any>([]);


watch(props.selectedUser, async (val) => {
  getMessages();


})
getMessages();
const handleTyping = (e) => {
 

  sendMessage(null, "typing");


}
</script>

<template>
  <div class="h-max-[200px] flex h-full flex-col bg-background mt-4 md:mt-0 md:p-4">
    <!-- <div class="flex items-center p-2">

      <div class="ml-auto flex items-center gap-2">
        
      </div>
     
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
    <Separator /> -->
    <Card ref="card" class="h-max-[200px] overflow-scroll overflow-x-hidden">
      <CardHeader class="flex flex-row items-center justify-between">
        <div class="flex items-center space-x-4">
          <Avatar>
            <AvatarImage src="/avatars/01.png" alt="Image" />
            <AvatarFallback>OM</AvatarFallback>
          </Avatar>
          <div>
            <p class="text-sm font-medium leading-none">{{ selectedUser.name }}</p>

          </div>
        </div>

      </CardHeader>
      <CardContent>
     
        <div class="space-y-4 p-5 ">

          <div v-for="(message, index) in messages" :key="index" :class="cn(
            'flex w-max max-w-[75%] flex-col gap-2 rounded-lg px-3 py-2 text-sm',
            message.role === 'user'
              ? 'ml-auto bg-primary text-primary-foreground'
              : 'bg-muted'
          )
            ">
            <p v-if="message.type != 'typing'">
              {{ message.content }}

            </p>
          
          </div>
          <p v-if=showTyping>
            <div class="chat-bubble">
              <div class="typing">
                <div class="dot"></div>
                <div class="dot"></div>
                <div class="dot"></div>
              </div>
            </div>
            </p>
        </div>
      </CardContent>
      <CardFooter>
        
      </CardFooter>
    </Card>
    <form class="flex w-full items-center space-x-2 mt-2" @submit.prevent="sendMessage">
          <Input @input="handleTyping" v-model="input" placeholder="Type a message..." class="flex-1" />

          <!-- <EmojiPicker picker-type="input" :native="true" @select="onSelectEmoji" /> -->
          <Button class="p-2.5 flex items-center justify-center" :disabled="inputLength === 0">
            <Send class="w-4 h-4" />
            <span class="sr-only">Send</span>
          </Button>
        </form>

  </div>
</template>



<style>
.chat-bubble {
  background-color: #E6F8F1;
  padding: 16px 28px;
  -webkit-border-radius: 20px;
  -webkit-border-bottom-left-radius: 2px;
  -moz-border-radius: 20px;
  -moz-border-radius-bottomleft: 2px;
  border-radius: 20px;
  border-bottom-left-radius: 2px;
  display: inline-block;
}

.typing {
  align-items: center;
  display: flex;
  height: 17px;
}

.typing .dot {
  animation: mercuryTypingAnimation 1.8s infinite ease-in-out;
  background-color: #6CAD96;
  border-radius: 50%;
  height: 7px;
  margin-right: 4px;
  vertical-align: middle;
  width: 7px;
  display: inline-block;
}

.typing .dot:nth-child(1) {
  animation-delay: 200ms;
}

.typing .dot:nth-child(2) {
  animation-delay: 300ms;
}

.typing .dot:nth-child(3) {
  animation-delay: 400ms;
}

.typing .dot:last-child {
  margin-right: 0;
}

@keyframes mercuryTypingAnimation {
  0% {
    transform: translateY(0px);
    background-color: #6CAD96;
  }

  28% {
    transform: translateY(-7px);
    background-color: #9ECAB9;
  }

  44% {
    transform: translateY(0px);
    background-color: #B5D9CB;
  }
}
</style>