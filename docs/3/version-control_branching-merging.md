# Branching & Merging

Before we get into the specifics of how version control works lets examine two of the most import concepts, branching and merging.

<p style="text-align:center">
<img src="img/git-branches.svg" alt="Branches" width="600px">
</p>

**Branching** allows engineers to work in isolation without disrupting the primary codebase. As multiple engineers work in parallel, each will make changes in their own branch independent of one another. Eventually they will want to integrate their code into the mainline. Branches should be short-lived and tied to work items.

**Merging** is the practice of reconciling changes between version controlled files from one branch to another branch. General merges can be performed effortlessly. Merging should happen as frequently as possible to provide maximum flow, faster feedback, and minimize conflicts.

> **What if there was no branching?** Everybody would be editing the live code, half-baked changes would bork the system, people would be stepping all over each other. And so we give individuals the illusion of frozen time, that they are the only ones changing the system and those changes can wait until they are fully baked before risking the system. But this is an illusion and eventually the price for it comes due. Who pays? When? How much? That's what these patterns are discussing: alternatives for paying the piper.
>
> _- Kent Beck, from [Martin Fowler's Patterns for Managing Source Code Branches](https://martinfowler.com/articles/branching-patterns.html)_

## Branching Strategies

As software development practices have evolved over time so have opinions about branching strategies. Sometimes these opinions can be a contentious issue. Read [GitFlow considered harmful](http://endoflineblog.com/gitflow-considered-harmful) and [the follow up](http://endoflineblog.com/follow-up-to-gitflow-considered-harmful) for some background on opinions of different strategies.


### Branching Best Practices

Now lets take a look at Liatio's point of view and branching.

**Guidelines for a good branching strategy**

- easy to implement and understand

- minimizes conflicts

- allows for testing and feedback to maintain code quality

- removes bottlenecks and unnecessary overhead

- makes the history of changes understandable and rolling back changes easier

---
<img src="img/git-icon-branch.svg" class="img-left" alt="Branch" height="80px" />

**When to create a branch?** <br>
A branch should be created only if required. This is most applicable when delivery teams have a history of creating long-lived release branches, whose changes often do not get merged into master.

---
<img src="img/git-icon-merge.svg" class="img-left" alt="Branch" height="80px" />

**When to merge changes?** <br>
Merging should happen as frequently as possible. As teams improve on creating smaller stories and engineers develop a clear way of working around pull requests, developers can merge code multiple times per day.
---

By eliminating long lived branches and merging smaller changes directly into a main branch teams can reduce merge conflicts and release code quicker. This approach is known as **truck-based development**. To achieve this kind of fast paced development, without sacrificing quality, it is important to have confidence in the quality of code being committed. This can be accomplished by including automated testing in the continuous integration process. We will cover automation and testing in more depth later in this chapter.