import { useEffect, useRef } from "react";
import { formatTime } from "../../lib/utils";
import { ChatMessage } from "@/lib/api";

export function ChatMessages({
  messages,
  userID,
}: {
  messages: ChatMessage[];
  userID: string;
}) {
  const bottomRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  return (
    <div className="flex-1 overflow-y-auto p-6 space-y-4 scrollbar-hide bg-white w-full">
      {messages &&
        messages.length > 0 &&
        messages.map((m) => {
          const isMine = m.userId === userID;
          return (
            <div
              key={m.id}
              className={`gap-2 animate-fade-in ${
                isMine ? "text-right" : "text-left"
              }`}
            >
              <div className="items-baseline gap-2">
                <span className="font-semibold text-gray-700 text-sm">
                  {m.username}
                </span>
                <span className="text-xs text-gray-400 ml-2">
                  {formatTime(new Date(m.timestamp))}
                </span>
              </div>

              <div className="inline-block rounded-2xl px-4 py-2 bg-amber-100 max-w-80 wrap-break-word whitespace-pre-wrap">
                {m.text}
              </div>
            </div>
          );
        })}

      <div ref={bottomRef} />
    </div>
  );
}
