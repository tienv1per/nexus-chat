import { redirect } from "next/navigation";

import { getInitialChatData } from "@/lib/api";

export default async function ChatIndexPage() {
  const data = await getInitialChatData();
  redirect(`/chat/${data.activeConversationId}`);
}
