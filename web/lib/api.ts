import { buildInitialData, seedUsers } from "@/lib/mock-data";
import type { ChatInitialData, ConversationID, User } from "@/lib/types";

function wait(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export async function getSeedUsers(): Promise<User[]> {
  await wait(80);
  return seedUsers;
}

export async function getInitialChatData(conversationId?: ConversationID): Promise<ChatInitialData> {
  await wait(80);
  return buildInitialData(conversationId);
}

export async function loginSeedUser(userId: string, deviceId: string) {
  await wait(220);
  const user = seedUsers.find((seedUser) => seedUser.id === userId);

  if (!user) {
    throw new Error("Seed user not found");
  }

  return {
    data: {
      user,
      token: `local.${user.username}.${deviceId}`
    }
  };
}
