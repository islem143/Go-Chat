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
import { useIntersectionObserver } from "@vueuse/core";


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
const card = useTemplateRef('card');











const inputLength = computed(() => input.value.trim().length);

let socket = new WebSocket("ws://localhost:8000/ws/private-chat/4");
socket.addEventListener("open", (event) => {

});
const props = defineProps(["selectedUser"]);



setTimeout(() => {
  card.value?.$el.scrollBy(0, 10000);

}, 100)
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
    setTimeout(() => {
      card.value?.$el.scrollBy(0, 10000);

    }, 0)
  }



};

socket.onmessage = (msg) => {
  let res = JSON.parse(msg.data);
  if (res.type == "typing") {
    showTyping.value = true;
    card.value?.$el.scrollBy(0, 10000);
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => {

      showTyping.value = false;
    }, 300);
  } else {
    
    messages.value.push({ role: "other", content: res.text, type: 'message' });
    setTimeout(() => {
      card.value?.$el.scrollBy(0, 10000);

    }, 0)
  }


};

const getMessages = async () => {
  let res = await Message.getChatMessages({ receiverId: props.selectedUser.id });


  res = res.map(p => ({...p, "content": p.text, role: props.selectedUser.id == p.receiver_id ? 'user' : 'other' }));


  messages.value = res;
};

let messages = ref<any>([]);

const messagesRefs = useTemplateRef('messagesRefs');


watch(props, () => {
  
getMessages();

});  
watch(messagesRefs, () => {
  if (messagesRefs) {
    setTimeout(()=>{
      const target = messagesRefs.value[messagesRefs.value?.length - 1];
      let lastMessage = messages.value[messages.value?.length - 1];
      if(lastMessage.role == 'other'){
        
      }
   
    
    
    const { stop } = useIntersectionObserver(
      target,
      ([{ isIntersecting }], observerElement) => {
        if (isIntersecting) {
          
          let lastMessage = messages.value[messages.value?.length - 1];
          
        
          
          
          
          Message.markAsRead({ messageId: lastMessage.id,receiverId: props.selectedUser.id });
        }
      },
    )
    },100)
    


  }

})
// onMounted(() => {

//   setTimeout(() => {


//   }, 100)
// })
getMessages();
const handleTyping = (e) => {

  sendMessage(null, "typing");


}
</script>

<template>

  <div class="h-max-[200px]  flex h-full flex-col mt-4 md:mt-0 md:p-4">
    <div class="flex items-center justify-between mb-4">
      <div class="flex items-center space-x-3">
        <Avatar class="h-10 w-10">
          <AvatarImage src="/avatars/01.png" alt="User avatar" />
          <AvatarFallback>{{ selectedUser.name[0] }}</AvatarFallback>
        </Avatar>
        <div>
          <p class="text-sm font-semibold">{{ selectedUser.name }}</p>
          <span class="text-xs text-gray-500">Active now</span>
        </div>
      </div>
      <div class="flex items-center space-x-2">
        <Button variant="ghost" size="icon" class="rounded-full">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M15 10h4.5a2.5 2.5 0 0 1 0 5H16"/><path d="M9 12h4"/><path d="M4 8v8"/></svg>
        </Button>
        <Button variant="ghost" size="icon" class="rounded-full">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/></svg>
        </Button>
      </div>
    </div>

    <Card ref="card" class="flex-1 h-[calc(100vh-250px)] overflow-y-scroll bg-gray-50">
      <CardContent class="p-4">
        <div v-if="messages.length > 0" class="space-y-4">
          <div v-for="(message, index) in messages" 
               ref="messagesRefs"
               :data-id="message.id"
               :key="index"
               :class="cn(
                 'flex w-max max-w-[75%] flex-col gap-1 rounded-2xl px-4 py-2 text-sm',
                 message.role === 'user'
                   ? 'ml-auto bg-blue-500 text-white'
                   : 'bg-white shadow-sm'
               )">
            <p v-if="message.type != 'typing'" class="break-words">
              {{ message.content }}
              <span v-if="message.role=='user'" class="flex justify-end gap-1 text-xs mt-1" :class="message.role === 'user' ? 'text-blue-100' : 'text-gray-400'">
                <span v-if="message.status === 'unread'">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3"><polyline points="20 6 9 17 4 12"/></svg>
                </span>
                <span v-if="message.status === 'read'" class="flex">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3"><path d="M18 6L7 17L2 12"/><path d="M22 10L13 19L11 17"/></svg>
                </span>
              </span>
            </p>
          </div>

          <div v-if="showTyping" 
               class="flex w-max max-w-[75%] flex-col gap-2 rounded-2xl px-3 py-2 bg-blue-400 shadow-sm">
            <div class="flex space-x-1">
              <div class="w-2 h-2 rounded-full bg-white animate-bounce "></div>
              <div class="w-2 h-2 rounded-full bg-white animate-bounce  [animation-delay:0.2s]"></div>
              <div class="w-2 h-2 rounded-full bg-white animate-bounce  [animation-delay:0.4s]"></div>
            </div>
          </div>

          
        </div>
        
      </CardContent>
    </Card>

    <form @submit.prevent="sendMessage" class="flex items-center gap-2 mt-4">
      <Input
        v-model="input"
        @input="handleTyping"
        placeholder="Type your message..."
        class="flex-1 bg-gray-50 border-gray-200 focus:ring-blue-500"
      />
      <Button 
        type="submit"
        :disabled="inputLength === 0"
        class="bg-blue-500 hover:bg-blue-600 text-white rounded-full p-2.5"
      >
        <Send class="w-5 h-5" />
        <span class="sr-only">Send message</span>
      </Button>
    </form>
  </div>
  <!-- <div class="h-max-[200px] flex h-full flex-col  mt-4 md:mt-0 md:p-4">
    
  
    <div class="flex items-center space-x-4 mb-4">
      <Avatar>
        <AvatarImage src="/avatars/01.png" alt="Image" />
        <AvatarFallback>CN</AvatarFallback>
      </Avatar>
      <div>
        <p class="text-sm font-medium leading-none">{{ selectedUser.name }}</p>

      </div>
    </div>
    <Card ref="card" class="xxx h-max-[200px] overflow-y-scroll overflow-x-hidden">

      <CardContent>

        <div v-if="messages.length > 0" class="space-y-4 p-5 ">

          <div v-for="(message, index) in messages" ref="messagesRefs" :data-id="message.id" :key="index" :class="cn(
            'flex w-max max-w-[75%] flex-col gap-2 rounded-lg px-3 py-2 text-sm',
            message.role === 'user'
              ? 'ml-auto bg-blue-600 text-primary-foreground'
              : 'bg-slate-200'
          )
            ">
            <p  v-if="message.type != 'typing'">
              {{ message.content }}
              <div class="flex justify-end gap-1 text-xs">
              <span v-if="message.status === 'unread'" class="text-gray-400">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3"><polyline points="20 6 9 17 4 12"/></svg>
              </span>
          
              <divc  v-if="message.status === 'read'" class=" flex gap-0">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3"><path d="M18 6L7 17L2 12"/><path d="M22 10L13 19L11 17"/></svg>

              </divc>
            </div>

            </p>
            
          </div>

          <div v-if="showTyping"
            class="flex w-max max-w-[75%] flex-col gap-2 rounded-lg px-3 py-2 text-sm  chat-bubble">
            <div class="typing">
              <div class="dot"></div>
              <div class="dot"></div>
              <div class="dot"></div>
            </div>
          </div>

        </div>
      </CardContent>
      <CardFooter>

      </CardFooter>
    </Card>
    <form class="flex w-full items-center space-x-2 mt-2" @submit.prevent="sendMessage">
      <Input @input="handleTyping" v-model="input" placeholder="Type a message..." class="flex-1" />

      <!-- <EmojiPicker picker-type="input" :native="true" @select="onSelectEmoji" /> -->
   

</template>



<style>
.chat-bubble {
  background-color: #008cff;
  padding: 8px 8px;
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
  height: 10px;
}

.typing .dot {
  animation: mercuryTypingAnimation 1.8s infinite ease-in-out;
  background-color: #ffffff;
  border-radius: 50%;
  height: 5px;
  margin-right: 4px;
  vertical-align: middle;
  width: 5px;
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

  }

  28% {
    transform: translateY(-7px);

  }

  44% {
    transform: translateY(0px);

  }
}
</style>