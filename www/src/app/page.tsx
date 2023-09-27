import "./page.css";
import Image from "next/image";
import Link from "next/link";
export default function Home() {
  return (
    <div className="main">
      <div className="hero">
        <div className="herocontent1">
          <div className="heroimage">
            <Image
              alt="hero image"
              src="/hero.png"
              layout="responsive"
              width={300}
              height={300}
            />
          </div>
          <div className="herotitle">
            <h3 className="heroc1h3">Kadrion Testops</h3>
            <p className="heroc1p">
              A set of Tools for perfect Continuous Testing in a <br /> reliable
              efficient way{" "}
            </p>
          </div>
        </div>
        <div className="herocontent2">
          <p className="heroc2p">
            Come Share your feedback or just hangout to talk more about this
            project
          </p>
          <button className="discord">
            <Link href="https://discord.gg/vmczkDud4h" target="_blank">
              Join The Discord
            </Link>
          </button>
        </div>
      </div>
      <div className="construction">
        <div className="inprogressimage">
          <Image
            alt="construction sign image"
            src="/construction.png"
            width={400}
            height={400}
          />
        </div>
        <div className="constructionmessage">
          The project is currently under active development! It will out for
          colloboration to the public soon. Stay tuned for future updates.
        </div>
      </div>
    </div>
  );
}
