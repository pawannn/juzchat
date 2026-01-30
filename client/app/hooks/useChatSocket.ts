import { useEffect, useRef, useState } from "react";
import { ChatMessage, fetchChats } from "@/lib/api";

export function useChatSocket(username: string) {
  const [messages, setMessages] = useState<ChatMessage[]>([]);
  const wsRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    const fetchAvailableChats = async () => {
      try {
        const chats = await fetchChats();
        setMessages(chats);
      } catch (err) {
        console.error("Failed to fetch chats", err);
      }
    };

    fetchAvailableChats();
  }, []);

  useEffect(() => {
    if (!username) return;

    const ws = new WebSocket(
      `${process.env.NEXT_PUBLIC_WEBSOCKET_BASE_URL}/chat`,
    );

    wsRef.current = ws;

    ws.onmessage = (e) => {
      const msg = JSON.parse(e.data);
      setMessages((prev) => [
        ...prev,
        {
          id: msg.id ?? crypto.randomUUID(),
          userId: msg.userId,
          username: msg.username,
          text: msg.text,
          timestamp: msg.timestamp * 1000,
        },
      ]);
    };

    return () => ws.close();
  }, [username]);

  const sendMessage = (payload: ChatMessage) => {
    if (wsRef.current?.readyState === WebSocket.OPEN) {
      wsRef.current.send(JSON.stringify(payload));
    }
  };

  return { messages, sendMessage };
}
