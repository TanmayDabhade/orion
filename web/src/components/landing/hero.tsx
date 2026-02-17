"use client"

import { Button } from "@/components/ui/button"
import { ArrowRight, Copy, Check } from "lucide-react"
import { motion, AnimatePresence } from "framer-motion"
import { useState, useEffect, useCallback } from "react"

const terminalLines = [
    { cmd: "o slack", output: "→ Opening Slack.app", delay: 800 },
    { cmd: "o d2l", output: "→ Opening https://d2l.msu.edu", delay: 600 },
    { cmd: "o code .", output: "→ Opening VS Code in ~/projects/orion", delay: 700 },
    { cmd: "o mail", output: "→ Opening https://mail.google.com", delay: 500 },
    { cmd: "o list --apps", output: "→ 127 apps detected", delay: 900 },
]

function TypingTerminal() {
    const [lines, setLines] = useState<{ text: string; type: "cmd" | "output" }[]>([])
    const [currentCmd, setCurrentCmd] = useState("")
    const [lineIndex, setLineIndex] = useState(0)

    const typeCommand = useCallback(async (cmd: string, output: string, delay: number) => {
        setCurrentCmd("")

        // Type each character
        for (let i = 0; i <= cmd.length; i++) {
            await new Promise(r => setTimeout(r, 50 + Math.random() * 40))
            setCurrentCmd(cmd.slice(0, i))
        }

        // Wait then show output
        await new Promise(r => setTimeout(r, 300))
        setLines(prev => [...prev, { text: `$ ${cmd}`, type: "cmd" }])
        setCurrentCmd("")

        await new Promise(r => setTimeout(r, 150))
        setLines(prev => [...prev, { text: output, type: "output" }])

        await new Promise(r => setTimeout(r, delay))
    }, [])

    useEffect(() => {
        let cancelled = false

        const run = async () => {
            while (!cancelled) {
                for (let i = 0; i < terminalLines.length; i++) {
                    if (cancelled) return
                    const { cmd, output, delay } = terminalLines[i]
                    setLineIndex(i)
                    await typeCommand(cmd, output, delay)
                }
                if (!cancelled) {
                    await new Promise(r => setTimeout(r, 1500))
                    setLines([])
                }
            }
        }

        run()
        return () => { cancelled = true }
    }, [typeCommand])

    return (
        <div className="relative rounded-2xl border border-white/[0.08] bg-[#050505] shadow-2xl overflow-hidden gradient-border">
            {/* Title bar */}
            <div className="flex items-center px-4 py-2.5 border-b border-white/[0.06] bg-white/[0.02]">
                <div className="flex gap-2">
                    <div className="w-3 h-3 rounded-full bg-white/20" />
                    <div className="w-3 h-3 rounded-full bg-white/15" />
                    <div className="w-3 h-3 rounded-full bg-white/10" />
                </div>
                <div className="flex-1 text-center">
                    <span className="text-[11px] text-white/20 font-mono">orion — zsh</span>
                </div>
            </div>

            {/* Terminal body */}
            <div className="p-5 font-mono text-[13px] leading-relaxed min-h-[280px] max-h-[280px] overflow-hidden">
                <AnimatePresence mode="sync">
                    {lines.map((line, i) => (
                        <motion.div
                            key={`${lineIndex}-${i}`}
                            initial={{ opacity: 0, y: 5 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ duration: 0.15 }}
                            className={line.type === "cmd" ? "text-white/60" : "text-white/30 ml-1 mb-2"}
                        >
                            {line.text}
                        </motion.div>
                    ))}
                </AnimatePresence>

                {/* Current typing line */}
                <div className="flex items-center text-white/60">
                    <span className="text-white/20">$ </span>
                    <span>{currentCmd}</span>
                    <span className="cursor-blink text-white/80 ml-0.5 font-light">▎</span>
                </div>
            </div>
        </div>
    )
}

export function Hero() {
    const [copied, setCopied] = useState(false)
    const installCmd = "curl -fsSL https://github.com/TanmayDabhade/orion/releases/latest/download/install.sh | sh"

    const copyToClipboard = () => {
        navigator.clipboard.writeText(installCmd)
        setCopied(true)
        setTimeout(() => setCopied(false), 2000)
    }

    return (
        <section className="relative min-h-screen flex flex-col justify-center overflow-hidden noise-bg">
            {/* Background effects — monochrome */}
            <div className="absolute inset-0 pointer-events-none">
                <div className="absolute top-1/3 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[600px] bg-white/[0.03] blur-[150px] rounded-full" />
            </div>

            <div className="max-w-6xl mx-auto px-6 w-full pt-24 pb-20 z-10">
                <div className="grid lg:grid-cols-2 gap-16 items-center">
                    {/* Left: Copy */}
                    <div className="space-y-8">
                        <motion.div
                            initial={{ opacity: 0, y: 20 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ duration: 0.5 }}
                        >
                            <div className="inline-flex items-center gap-2 px-3 py-1 rounded-full border border-white/[0.08] bg-white/[0.03] text-[12px] text-white/40 mb-6">
                                <span className="flex h-1.5 w-1.5 rounded-full bg-white/60 animate-pulse" />
                                v0.7.0 Available
                            </div>
                        </motion.div>

                        <motion.h1
                            initial={{ opacity: 0, y: 20 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ duration: 0.5, delay: 0.1 }}
                            className="text-5xl lg:text-6xl font-bold tracking-tight leading-[1.1]"
                        >
                            <span className="text-white">Your terminal,{" "}</span>
                            <br />
                            <span className="text-white/60">supercharged.</span>
                        </motion.h1>

                        <motion.p
                            initial={{ opacity: 0, y: 20 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ duration: 0.5, delay: 0.2 }}
                            className="text-lg text-white/35 leading-relaxed max-w-lg"
                        >
                            The ultra-fast CLI launcher for macOS. Open any app, manage URL shortcuts, and automate your daily workflow — all from a single command.
                        </motion.p>

                        <motion.div
                            initial={{ opacity: 0, y: 20 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ duration: 0.5, delay: 0.3 }}
                            className="flex flex-col sm:flex-row gap-3"
                        >
                            <Button
                                size="lg"
                                className="h-11 px-6 bg-white text-black hover:bg-white/90 font-medium text-sm rounded-xl relative overflow-hidden shine-hover font-mono"
                                onClick={copyToClipboard}
                            >
                                <div className="flex items-center gap-2">
                                    <span className="text-black/40">$</span>
                                    <span className="truncate max-w-[260px]">curl -fsSL .../install.sh | sh</span>
                                </div>
                                {copied ? (
                                    <Check className="ml-2 h-3.5 w-3.5 text-black shrink-0" />
                                ) : (
                                    <Copy className="ml-2 h-3.5 w-3.5 opacity-40 shrink-0" />
                                )}
                            </Button>

                            <Button
                                variant="outline"
                                size="lg"
                                className="h-11 px-6 bg-transparent border-white/[0.08] hover:bg-white/[0.04] hover:border-white/[0.12] text-white/50 text-sm rounded-xl"
                                asChild
                            >
                                <a href="https://github.com/TanmayDabhade/orion" target="_blank">
                                    View on GitHub
                                </a>
                            </Button>
                        </motion.div>

                        <motion.div
                            initial={{ opacity: 0 }}
                            animate={{ opacity: 1 }}
                            transition={{ duration: 0.5, delay: 0.5 }}
                            className="flex items-center gap-6 pt-2 text-[13px] text-white/25"
                        >
                            <span className="flex items-center gap-1.5">
                                <span className="h-1 w-1 rounded-full bg-white/40" />
                                macOS native
                            </span>
                            <span className="flex items-center gap-1.5">
                                <span className="h-1 w-1 rounded-full bg-white/40" />
                                Built in Go
                            </span>
                            <span className="flex items-center gap-1.5">
                                <span className="h-1 w-1 rounded-full bg-white/40" />
                                Open source
                            </span>
                        </motion.div>
                    </div>

                    {/* Right: Terminal */}
                    <motion.div
                        initial={{ opacity: 0, y: 30, scale: 0.98 }}
                        animate={{ opacity: 1, y: 0, scale: 1 }}
                        transition={{ duration: 0.7, delay: 0.4 }}
                        className="hidden lg:block"
                    >
                        <div className="glow rounded-2xl">
                            <TypingTerminal />
                        </div>
                    </motion.div>
                </div>
            </div>
        </section>
    )
}
