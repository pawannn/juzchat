"use client";

import { useUsername } from "./hooks/useUsername";
import { useChatSocket } from "./hooks/useChatSocket";
import { useNow } from "./hooks/useNow";
import { NamePrompt } from "./components/NamePrompt";
import { ChatMessages } from "./components/ChatMessages";
import { ChatInput } from "./components/ChatInput";
import { ChatHeader } from "./components/ChatHeader";

export default function ChatPage() {
  useNow();

  const { username, userID, showPrompt, saveName } = useUsername();
  const { messages, sendMessage } = useChatSocket(username);

  if (showPrompt) return <NamePrompt onSubmit={saveName} />;

  return (
    <div className="flex md:items-center md:justify-center h-11/12 md:h-screen bg-linear-to-br from-amber-50 via-amber-25 to-green-50">
      <div className="w-full max-w-2xl flex flex-col h-svh md:h-150 bg-white rounded-3xl shadow-lg overflow-hidden">
        <ChatHeader />
        <ChatMessages messages={messages} userID={userID} />
        <ChatInput
          onSend={(text) =>
            sendMessage({
              id: crypto.randomUUID(),
              userId: userID,
              username,
              text,
              timestamp: Date.now(),
            })
          }
        />
      </div>
    </div>
  );
}
