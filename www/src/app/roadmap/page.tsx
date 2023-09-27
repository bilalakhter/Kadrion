"use client";
import {
  VerticalTimeline,
  VerticalTimelineElement,
} from "react-vertical-timeline-component";
import "react-vertical-timeline-component/style.min.css";
import "./page.css";
import Image from "next/image";
export default function Roadmap() {
  return (
    <div className="roadmap-container">
      <VerticalTimeline>
        <VerticalTimelineElement
          className="vertical-timeline-element--work"
          contentStyle={{ background: "#516BAA", color: "#fff" }}
          contentArrowStyle={{ borderRight: "7px solid  #516BAA" }}
          iconStyle={{
            background: "#516BAA",
            color: "#fff",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
          }}
          icon={
            <Image
              src="/flag.png"
              width={50}
              height={50}
              alt="Picture of the author"
            />
          }
        >
          <h3 className="vertical-timeline-element-title">
            Project Initialization
          </h3>
          <p>
            To establish the foundational structure and environment for our
            project. With a well-defined directory structure and initial
            documentation, we set the stage for seamless collaboration.
          </p>
        </VerticalTimelineElement>
        <VerticalTimelineElement
          className="vertical-timeline-element--work"
          contentStyle={{ background: "#516BAA", color: "#fff" }}
          contentArrowStyle={{ borderRight: "7px solid  #516BAA" }}
          iconStyle={{
            background: "#516BAA",
            color: "#fff",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
          }}
          icon={
            <Image
              src="/gears.png"
              width={50}
              height={50}
              alt="Picture of the author"
            />
          }
        >
          <h3 className="vertical-timeline-element-title">
            Core Feature Implementation
          </h3>
          <p>
            To build functionalities that define the essence of project.With
            Core library completed this shows significant leap forward.
          </p>
        </VerticalTimelineElement>
        <VerticalTimelineElement
          className="vertical-timeline-element--work"
          contentStyle={{ background: "#516BAA", color: "#fff" }}
          contentArrowStyle={{ borderRight: "7px solid  #516BAA" }}
          iconStyle={{
            background: "#516BAA",
            color: "#fff",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
          }}
          icon={
            <Image
              src="/magnifyglass.png"
              width={50}
              height={50}
              alt="Picture of the author"
            />
          }
        >
          <h3 className="vertical-timeline-element-title">
            Testing and Refinement
          </h3>
          <p>
            The core Library is testing on multiple scenario ,THis helps to
            furthur fine tune the project leading it to mature project for
            colloboration.
          </p>
        </VerticalTimelineElement>
        <VerticalTimelineElement
          className="vertical-timeline-element--work"
          contentStyle={{ background: "#516BAA", color: "#fff" }}
          contentArrowStyle={{ borderRight: "7px solid  #516BAA" }}
          iconStyle={{
            background: "#516BAA",
            color: "#fff",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
          }}
          icon={
            <Image
              src="/announcement.png"
              width={50}
              height={50}
              alt="Picture of the author"
            />
          }
        >
          <h3 className="vertical-timeline-element-title">
            Open for Collaboration
          </h3>
          <p>
            At this stage the project will move towards public driven open
            source development
          </p>
        </VerticalTimelineElement>
      </VerticalTimeline>
    </div>
  );
}
