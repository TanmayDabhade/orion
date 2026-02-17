import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Orion - Your Terminal, Supercharged.",
  description: "The ultra-fast CLI launcher for macOS. Open apps, manage shortcuts, and navigate your system at the speed of thought. Built in Go.",
  keywords: ["CLI", "Launcher", "macOS", "Productivity", "Terminal", "Go"],
  authors: [{ name: "Tanmay Dabhade" }],
  openGraph: {
    title: "Orion - Your Terminal, Supercharged.",
    description: "The ultra-fast CLI launcher for macOS. Open apps, manage shortcuts, and navigate your system at the speed of thought.",
    type: "website",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark" style={{ colorScheme: "dark" }}>
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased bg-black text-white selection:bg-white/20`}
      >
        {children}
      </body>
    </html>
  );
}
