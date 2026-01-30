import { useState } from "react";

const USER_ID_KEY = "juzchat:userId";
const USERNAME_KEY = "juzchat:username";

export const useUsername = () => {
  const [userID] = useState(() => {
    if (typeof window === "undefined") return "";

    let uid = localStorage.getItem(USER_ID_KEY);
    if (!uid) {
      uid = crypto.randomUUID();
      localStorage.setItem(USER_ID_KEY, uid);
    }
    return uid;
  });

  const [username, setUsername] = useState(() => {
    if (typeof window === "undefined") return "";
    return localStorage.getItem(USERNAME_KEY) ?? "";
  });

  const [showPrompt, setShowPrompt] = useState(() => {
    if (typeof window === "undefined") return false;
    return !localStorage.getItem(USERNAME_KEY);
  });

  const saveName = (name: string) => {
    localStorage.setItem(USERNAME_KEY, name);
    setUsername(name);
    setShowPrompt(false);
  };

  return { username, userID, showPrompt, saveName };
};
