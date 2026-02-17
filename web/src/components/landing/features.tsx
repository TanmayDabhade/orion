"use client"

import { motion } from "framer-motion"
import { Zap, Command, Rocket, Keyboard, Globe, Shield } from "lucide-react"

const features = [
    {
        title: "Smart Shortcuts",
        description: "Create aliases for any URL or command. Type less, do more.",
        icon: Command,
        example: "o d2l → opens course page",
    },
    {
        title: "App Launcher",
        description: "Auto-detects every installed app. Zero configuration.",
        icon: Rocket,
        example: "o slack → opens Slack.app",
    },
    {
        title: "Instant Startup",
        description: "Native Go binary. No Electron. No wait time. Just results.",
        icon: Zap,
        example: "~2ms cold start",
    },
    {
        title: "Keyboard First",
        description: "Your hands stay on the keyboard. No mouse required.",
        icon: Keyboard,
        example: "Tab completion built-in",
    },
    {
        title: "URL Management",
        description: "Manage bookmarks from the terminal. Open any site instantly.",
        icon: Globe,
        example: "o mail → opens Gmail",
    },
    {
        title: "Safe by Default",
        description: "Conservative risk gating. Dangerous commands need confirmation.",
        icon: Shield,
        example: "Risk level: low ✓",
    },
]

export function Features() {
    return (
        <section id="features" className="relative py-28 overflow-hidden">
            <div className="max-w-6xl mx-auto px-6">
                {/* Section Header */}
                <motion.div
                    initial={{ opacity: 0, y: 20 }}
                    whileInView={{ opacity: 1, y: 0 }}
                    viewport={{ once: true }}
                    className="text-center mb-16"
                >
                    <p className="text-[13px] text-white/30 uppercase tracking-widest mb-3 font-medium">Features</p>
                    <h2 className="text-3xl md:text-4xl font-bold tracking-tight text-white mb-4">
                        Everything you need. Nothing you don&apos;t.
                    </h2>
                    <p className="text-white/35 text-base max-w-lg mx-auto">
                        Orion replaces Spotlight, Alfred, and bookmark managers with a single, fast command.
                    </p>
                </motion.div>

                {/* Feature Grid */}
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                    {features.map((feature, index) => (
                        <motion.div
                            key={index}
                            initial={{ opacity: 0, y: 15 }}
                            whileInView={{ opacity: 1, y: 0 }}
                            viewport={{ once: true }}
                            transition={{ delay: index * 0.05 }}
                            className="group relative p-6 rounded-2xl border border-white/[0.06] bg-white/[0.02] hover:bg-white/[0.04] hover:border-white/[0.1] transition-all duration-300"
                        >
                            <div className="h-9 w-9 rounded-lg bg-white/[0.06] flex items-center justify-center mb-4 group-hover:bg-white/[0.1] transition-colors">
                                <feature.icon className="h-4.5 w-4.5 text-white/60" />
                            </div>
                            <h3 className="text-[15px] font-semibold text-white mb-1.5">{feature.title}</h3>
                            <p className="text-[13px] text-white/35 leading-relaxed mb-3">{feature.description}</p>
                            <p className="text-[12px] font-mono text-white/20">{feature.example}</p>
                        </motion.div>
                    ))}
                </div>
            </div>
        </section>
    )
}
