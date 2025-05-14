"use client";

import type { Session } from "next-auth";
import { SessionProvider } from "next-auth/react";
import { ThemeProvider } from "next-themes";

export default function Providers({
  session,
  children,
}: {
  session: Session | null;
  children: React.ReactNode;
}) {
  return (
    <SessionProvider session={session}>
      {" "}
      <ThemeProvider attribute={"class"} defaultTheme="dark">
        {children}
      </ThemeProvider>
    </SessionProvider>
  );
}
