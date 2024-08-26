import Image from "next/image"
import Link from "next/link"
import { ChevronLeft, ChevronRight } from "lucide-react"

import { env } from "@/env.mjs"
import { siteConfig } from "@/config/site"
import { cn } from "@/lib/utils"
import { Button, buttonVariants } from "@/components/ui/button"
import { Icons } from "@/components/icons"

import AnimatedGradientText from "./components/animated-gradient"
import { MarqueeDemo } from "./components/testimonials"
import { services } from "./services"

export default async function IndexPage() {
  return (
    <>
      <section className="space-y-6 pb-8 pt-6 md:pb-12 md:pt-6 lg:py-20">
        <div className="container z-10 flex max-w-5xl flex-col items-center gap-4 text-center">
          <Link
            href={"/docs"}
            className="z-10 flex items-center justify-center"
          >
            <AnimatedGradientText>
              🎉 <hr className="mx-2 h-4 w-px shrink-0 bg-gray-300" />{" "}
              <span
                className={cn(
                  `animate-gradient inline bg-gradient-to-r from-[#ffaa40] via-[#9c40ff] to-[#ffaa40] bg-[length:var(--bg-size)_100%] bg-clip-text text-transparent`
                )}
              >
                Explore all benifits
              </span>
              <ChevronRight className="ml-1 size-3 transition-transform duration-300 ease-in-out group-hover:translate-x-0.5" />
            </AnimatedGradientText>
          </Link>
        </div>
        <div className="container flex max-w-7xl flex-col items-center gap-4 text-center">
          <h1 className="font-heading text-2xl sm:text-4xl md:text-5xl lg:text-6xl">
            Breaking Barriers, Building Trust: Your Source for Decentralized
            Journalism
          </h1>
          <p className="mt-6 flex gap-4 text-xl lg:gap-8 2xl:gap-12">
            Unleashing the Power of Truth: Decentralized News for a Connected
            World
          </p>
        </div>

        <div className="flex flex-wrap items-center justify-center gap-6 pt-12">
          <Link
            href={"/dashboard"}
            className={buttonVariants({
              className: "w-[300px] border-2 border-yellow-600 p-6",
              variant: "secondary",
            })}
          >
            Personalized News
          </Link>
          <Link
            href={"/programs#trending"}
            className={buttonVariants({
              className: "w-[300px] border-2 border-yellow-600 p-6",
              variant: "secondary",
            })}
          >
            Top Trending News
          </Link>
          <Link
            href={"/programs#research"}
            className={buttonVariants({
              className: "w-[300px] border-2 border-yellow-600 p-6",
              variant: "secondary",
            })}
          >
            Sports News
          </Link>
          <Link
            href={"/programs#upcoming"}
            className={buttonVariants({
              className: "w-[300px] border-2 border-yellow-600 p-6",
              variant: "secondary",
            })}
          >
            Financial News
          </Link>
        </div>
      </section>
      <section className="w-full bg-secondary py-8 md:py-12 lg:py-24">
        <h3 className="mb-12 text-center font-heading text-xl sm:text-2xl md:text-3xl lg:text-4xl">
          WHY WE&apos;RE OFFERING
        </h3>
        <div className="flex items-center justify-center gap-12">
          {[
            {
              title: "Collection",
              subtitle: "500+ ARTICLES",
              value: "by trusted journalists",
            },
            {
              title: "Fake news",
              subtitle: "90% catch rate",
              value: "By Our custom AI model",
            },
            {
              title: "Authenticity",
              subtitle: "Blockchain technology",
              value: "for all of our news articles",
            },
            {
              title: "AI/ML Integration",
              subtitle: "5 unique AI features",
              value: "Summary, Grammar checks etc.",
            },
          ].map((i) => (
            <div className="flex h-72 w-60 flex-col justify-between rounded-2xl bg-primary p-4 py-12 font-extrabold text-primary-foreground">
              <h4 className="text-center text-xl">{i.title}</h4>
              <h4 className="text-center text-4xl">{i.subtitle}</h4>
              <h4 className="text-center text-xl">{i.value}</h4>
            </div>
          ))}
        </div>
      </section>
      <section className="w-full bg-primary/90 p-12">
        <h3 className="mb-12 text-center font-heading text-xl text-primary-foreground sm:text-2xl md:text-3xl lg:text-4xl">
          Trending News
        </h3>
        <div className="flex flex-wrap items-center justify-center gap-24">
          {[
            {
              title: "50B  people dead in 2024",
              desc: "Lorem ipsum dolor, sit amet consectetur adipisicing elit. Maxime iusto sed, repellendus delectus, nostrum nihil fugiat laborum architecto, est eum doloremque inventore beatae distinctio eius error amet incidunt iure minima provident voluptatibus. Deleniti illo commodi amet. Suscipit possimus, aperiam maiores velit saepe animi, commodi sapiente error exercitationem nostrum magni voluptates!",
            },
            {
              title: "10,000 Suicides in India",
              desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequuntur quia saepe eligendi possimus nihil fuga optio ex aut sequi impedit, ipsa facere placeat odit, incidunt, harum accusantium velit ipsum sit dignissimos ut aspernatur rem soluta! Eligendi nihil commodi eos ex veniam modi voluptate dolorum quia aperiam repudiandae, explicabo alias. Dolore possimus laudantium, quasi libero aperiam incidunt laborum illo perspiciatis harum tempore doloribus, nulla molestiae dolorum maxime nam voluptas. Cumque, eaque!",
            },
          ].map((i) => (
            <div className="relative h-[500px] w-[450px] overflow-hidden rounded-3xl border bg-primary/70 pt-12 text-primary-foreground shadow-xl">
              <h3 className="mb-8 px-12 text-2xl">{i.title}</h3>
              <p className="px-12">Trending since 2 days</p>
              <p className="mb-6 px-12">⭐⭐⭐⭐⭐</p>
              <p className="px-12 text-xs">{i.desc}</p>
              <Link
                href={"/"}
                className="absolute bottom-0 flex w-full items-center justify-center gap-6 border-t bg-secondary/30 p-3 text-center text-2xl"
              >
                View Full Article
                <ChevronRight />
              </Link>
            </div>
          ))}
        </div>
        <h3 className="mb-12 mt-24 text-center font-heading text-xl text-primary-foreground sm:text-2xl md:text-3xl lg:text-4xl">
          Recent Events
        </h3>
        <div className="flex flex-wrap items-center justify-center gap-24">
          {[
            {
              title: "Lorem ipsum dolor sit amet consectetur adipisicing.",
              desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Itaque placeat ipsum quibusdam nisi excepturi commodi, soluta facere esse tempora sed illum aliquid ut alias facilis laudantium eum rerum asperiores dolorum quis, eaque dolores aperiam minima? Vero veritatis pariatur modi voluptatum quam sequi consectetur quae cum repellendus eligendi! Voluptatum ducimus possimus delectus soluta voluptas tempore tempora odio, error maiores nemo nam corporis obcaecati iste in facilis commodi facere, accusamus fugit sit dolores veniam dicta quidem? Fugit maxime corrupti aperiam itaque nulla. Ipsa laudantium, similique corrupti esse ",
            },
            {
              title: "Lorem ipsum dolor sit amet consectetur, adipisicing",
              desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Sunt beatae quae, consequuntur hic maiores, ullam, tempore blanditiis quam reprehenderit dolores animi cupiditate molestiae officiis porro eos. Nisi tempore exercitationem similique nostrum quo quis illo? Libero consequuntur blanditiis sequi repudiandae voluptas, consequatur vel. Nostrum consectetur quibusdam laboriosam! Obcaecati voluptatem ea quaerat sequi delectus porro sed, praesentium omnis amet quasi eum aspernatur vel. Atque debitis, excepturi quod, nobis deserunt voluptates esse vero sapiente explicabo consectetur tempore dolores inventore corporis at tenetur. Error.",
            },
          ].map((i) => (
            <div className="relative h-[550px] w-[450px] overflow-hidden rounded-3xl border bg-primary/70 pt-12 text-primary-foreground shadow-xl">
              <h3 className="mb-8 px-12 text-2xl">{i.title}</h3>
              <p className="px-12">Happened 10hr. ago</p>
              <p className="mb-2 px-12">⭐⭐⭐⭐⭐</p>
              <div className="mb-6 flex items-center gap-4 px-12">
                <span className="size-12 rounded-full bg-green-700" />
                <p className="text-xl">Registration Closed</p>
              </div>
              <p className="px-12 text-xs">{i.desc}</p>
              <Link
                href={"/"}
                className="absolute bottom-0 flex w-full items-center justify-center gap-6 border-t bg-secondary/30 p-3 text-center text-2xl"
              >
                View Full Article
                <ChevronRight />
              </Link>
            </div>
          ))}
        </div>
      </section>
      <section className="w-full bg-primary/90 py-12">
        <h3 className="mb-12 mt-24 pl-36 font-heading text-xl text-primary-foreground sm:text-2xl md:text-3xl lg:text-4xl">
          Fuse News advantage
        </h3>
        <div className="flex bg-background p-36 pt-12">
          <div className="">
            <h4 className="mb-6 text-3xl">Learn from the experts</h4>
            <p className="max-w-lg text-xl">
              Develop real-world skills with our news and projects designed by
              industry experts and practitioners
            </p>
          </div>
          <div className="flex"></div>
        </div>
        <div className="flex bg-background p-36 pt-12">
          <div className="">
            <h4 className="mb-6 text-3xl">Why Fuse?</h4>
            <p className="max-w-lg text-xl">
              We focus on personalized news feeds and real-world application
              with live projects to craft your X factor and help you uniquely
              stand out in your college applications
            </p>
          </div>
          <div className="flex"></div>
        </div>
        <div className="flex justify-center gap-16 bg-background p-36 pt-12">
          {[
            {
              title: "Real Hands-on reporters",
              desc: "Work on real world problems and solve using technology",
            },
            {
              title: "Real Hands-on reporters",
              desc: "Solve problems with expert guidance in the field",
            },
            {
              title: "Real Hands-on reporters",
              desc: "Receive signed letters of recommendation and certificate to showcase your contribution value",
            },
          ].map((i) => (
            <div className="h-52 w-80 rounded-3xl border bg-primary p-8 text-primary-foreground">
              <h3 className="mb-8 text-xl font-bold">{i.title}</h3>
              <h3>{i.desc}</h3>
            </div>
          ))}
        </div>
      </section>
      <section className="flex flex-col bg-primary text-primary-foreground">
        <h3 className="mb-6 text-center text-5xl font-bold">
          Why people love Fuse News?
        </h3>
        <p className="mb-12 text-center">
          SEE WHAT HIGH READERS HAVE TO SAY ABOUT OUR ARTICLES
        </p>
        <div className="mb-16">
          <MarqueeDemo />
        </div>
        <div className="mx-auto flex max-w-7xl flex-wrap items-center justify-around gap-12">
          {[
            {
              rating: "4.9 Rating",
              desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Facere, dolorum commodi reiciendis iste nam itaque impedit odit provident placeat animi, fugiat, hic voluptas repudiandae? Beatae dolore quidem ipsam sint. Ad consequuntur, fugit dolor magni quia dignissimos quibusdam autem cupiditate quam.",
              id: "Riya Sharma Texas",
            },
            {
              rating: "4.9 Rating",
              desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Facere, dolorum commodi reiciendis iste nam itaque impedit odit provident placeat animi, fugiat, hic voluptas repudiandae? Beatae dolore quidem ipsam sint. Ad consequuntur, fugit dolor magni quia dignissimos quibusdam autem cupiditate quam.",
              id: "Riya Sharma Texas",
            },
          ].map((i) => (
            <div className="h-[350px] w-[500px] rounded-3xl border p-8">
              <h2 className="mb-12 text-4xl font-extrabold">{i.rating}</h2>
              <p className="mb-10">{i.desc}</p>
              <p>{i.id}</p>
            </div>
          ))}
          <div className="my-36 flex items-center justify-center">
            <div className="flex w-[90%] items-center gap-8 rounded-3xl bg-blue-900 p-8 text-primary">
              <div>
                <h2 className="mb-12 text-4xl font-extrabold">
                  Still Confused?
                </h2>
                <h2 className="text-2xl font-extrabold">
                  Need more information about the articles? Talk to an expert
                  =&gt;
                </h2>
              </div>
              <Button className="bg-yellow-600 p-8 px-12 hover:bg-yellow-500">
                Contact Us <ChevronRight />
              </Button>
            </div>
          </div>
        </div>
      </section>
      <section
        id="open-source"
        className="w-full bg-primary py-8 md:py-12 lg:py-24"
      >
        <div className="flex h-[20vh] w-full flex-col">
          <h3 className="mx-auto mb-12 max-w-5xl text-center font-heading text-5xl text-primary-foreground">
            Boost your general knowledge with our tech & research news
          </h3>
          <Link
            href={"/"}
            className={buttonVariants({
              className:
                "rounded-[20px] bg-primary-foreground px-16 text-xl mx-auto text-white mb-36 hover:bg-secondary",
            })}
          >
            Explore All Articles
          </Link>
        </div>
      </section>
    </>
  )
}
