# Web App

Phase 2 implements the Next.js App Router chat UI here.

## Commands

```bash
npm install
npm run dev
npm test
npm run build
```

## Routes

- `/login`: dev login with seeded users and local device id.
- `/chat`: redirects to the first seeded conversation.
- `/chat/[conversationId]`: responsive chat workspace with mock socket state.
- `/dashboard`, `/admin`, `/settings`: lightweight shell destinations for Phase 2 navigation.

The planned screens are documented in:

- `docs/project_plans/implementation_plans/features/local-chat-system-v1-ui-screens.md`
