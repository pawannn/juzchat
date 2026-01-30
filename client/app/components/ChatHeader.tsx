"use client";

export const ChatHeader = () => {
  return (
    <div className="bg-linear-to-r from-amber-50 to-green-50 px-6 py-4 border-b border-amber-100">
      <h1 className="text-3xl font-bold text-gray-800 text-balance">JuzChat</h1>
      <p className="text-sm text-gray-500 mt-1">Just come. Just chat.</p>
      <p className="text-xs text-gray-400 mt-3 flex items-center gap-1">
        <span className="inline-block w-2 h-2 bg-green-400 rounded-full animate-pulse" />
        Messages disappear after 3 hours
      </p>
    </div>
  );
};
