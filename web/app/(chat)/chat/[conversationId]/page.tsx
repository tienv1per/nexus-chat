import { notFound } from "next/navigation";

import { ChatShell } from "@/components/chat/chat-shell";
import { getInitialChatData } from "@/lib/api";

type ConversationPageProps = {
  params: Promise<{ conversationId: string }>;
};

export default async function ConversationPage({ params }: ConversationPageProps) {
  const { conversationId } = await params;
  const data = await getInitialChatData(conversationId);

  if (!data.conversations.some((conversation) => conversation.id === conversationId)) {
    notFound();
  }

  return <ChatShell initialData={data} />;
}
