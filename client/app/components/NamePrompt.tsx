import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useState } from "react";

export const NamePrompt = ({ onSubmit }: { onSubmit: (n: string) => void }) => {
  const [value, setValue] = useState("");

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const next = e.target.value;

    if (/^[a-zA-Z0-9]*$/.test(next)) {
      setValue(next);
    }
  };

  const handleSubmit = () => {
    if (!value) return;
    onSubmit(value);
  };

  return (
    <div className="fixed inset-0 bg-black/20 backdrop-blur-sm flex items-center justify-center z-50">
      <div className="bg-white p-8 rounded-3xl shadow-xl w-full max-w-md">
        <h2 className="text-2xl font-bold text-center mb-6">
          What should we call you? 😊
        </h2>

        <Input
          autoFocus
          value={value}
          onChange={handleChange}
          onKeyDown={(e) => e.key === "Enter" && handleSubmit()}
          placeholder="Only letters and numbers"
        />

        <Button
          className="w-full mt-4"
          disabled={!value}
          onClick={handleSubmit}
        >
          Start chatting →
        </Button>
      </div>
    </div>
  );
};
