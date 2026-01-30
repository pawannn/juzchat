import axios from "axios";

export interface ChatMessage {
  id: string;
  username: string;
  text: string;
  timestamp: number;
  userId: string;
}

const backendBase = process.env.NEXT_PUBLIC_BACKEND_BASE_URL;

const api = axios.create({
  baseURL: backendBase,
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
    "ngrok-skip-browser-warning": "69420",
  },
});

export const fetchChats = async (): Promise<ChatMessage[]> => {
  const res = await api.get<ChatMessage[]>("/chats");
  return res.data;
};
