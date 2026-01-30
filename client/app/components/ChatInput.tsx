import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Send } from "lucide-react";
import { useState } from "react";

export const ChatInput = ({ onSend }: { onSend: (t: string) => void }) => {
  const [value, setValue] = useState("");

  const send = () => {
    if (!value.trim()) return;
    onSend(value);
    setValue("");
  };

  return (
    <div className="p-4 flex gap-2 bg-white">
      <Input
        value={value}
        onChange={(e) => setValue(e.target.value)}
        onKeyDown={(e) => e.key === "Enter" && send()}
        placeholder="Juz say something… 😊"
        className="flex-1 rounded-full border-amber-200 bg-amber-50 text-gray-800 placeholder:text-gray-400 focus:border-amber-300 focus:bg-white transition-colors"
      />
      <Button onClick={send}>
        <Send />
      </Button>
    </div>
  );
};
