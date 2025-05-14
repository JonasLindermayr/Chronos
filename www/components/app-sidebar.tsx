"use client";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import {
  History,
  LayoutDashboard,
  Settings,
  Users,
  UsersIcon,
} from "lucide-react";
import Link from "next/link";
import { Button } from "./ui/button";
import { signOut } from "next-auth/react";

const items = [
  {
    title: "Dashboard",
    url: "/app",
    icon: LayoutDashboard,
  },
  {
    title: "Employee's",
    url: "/app/employees",
    icon: Users,
  },
  {
    title: "Time-Logs",
    url: "/app/logs",
    icon: History,
  },
  {
    title: "Settings",
    url: "/app/settings",
    icon: Settings,
  },
];

export function AppSidebar() {
  return (
    <Sidebar>
      <SidebarHeader>
        <Link href="#" className="flex items-center gap-2 font-medium">
          <div className="flex h-6 w-6 items-center justify-center rounded-md bg-primary text-primary-foreground">
            <UsersIcon className="size-4" />
          </div>
          Chronos
        </Link>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Application</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <Link href={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <SidebarMenuButton asChild>
          <Button onClick={() => signOut()}>Logout</Button>
        </SidebarMenuButton>
        <SidebarMenuItem className="text-xs text-muted-foreground text-center">
          v{process.env.NEXT_PUBLIC_VERSION}
        </SidebarMenuItem>
      </SidebarFooter>
    </Sidebar>
  );
}
