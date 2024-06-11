# Prompt Cookbook



### Content Abstract

```
Please respond only in the english language. Do not explain what you are doing. Do not self reference. The response you provide is being used as input in a functioning application, so adherance to structure and schema is critical. Examine the text contents of this JSON file only. Ignore all code, scripting, markup, markdown and formatting.

URL: Provide the URL to the page, if supplied in the "url" field.
Title: Provide the title to the page, if supplied in the "title" field.
Meta Title: Provide the Meta Title to the page, if supplied in the "meta_title" field.
Meta Description: Provide the Meta Description, if provided in the "meta_description" field.

Abstract:
Summarize the content in less than 300 words.

Category: Which of the following general categories would best represent the content? Provide only one answer with no explanation or additional text: "News, Politics, Business, Technology, Science, Health, Environment, Education, Sports, Entertainment, Culture, History, Geography, Travel, Food & Drink, Lifestyle, Fashion, Art, Literature, Music, Movies & TV, Theater, Architecture, Photography, Philosophy, Religion, Sociology, Psychology, Economics, Finance, Law, Government, Military, Crime & Justice, Real Estate, Personal Development, Relationships, Parenting, Home & Garden, DIY & Crafts, Automobiles, Aviation, Space, Nature & Wildlife, Pets & Animals, Hobbies, Gaming, Fitness, Adventure, Wellness, Holidays & Celebrations, Folklore & Mythology, Digital Media, Social Media, Advertising, Marketing, Sales, Productivity, Career & Employment, Startups & Entrepreneurship, Innovation, Blockchain & Cryptocurrencies, Artificial Intelligence, Virtual Reality & Augmented Reality, Robotics, Space Exploration, Oceanography, Meteorology, Agriculture, Biotechnology, Public Health, Epidemics & Pandemics, Alternative Medicine, Mental Health, Addiction & Recovery, Disaster Preparedness, Urban Planning, Transportation, Infrastructure, Human Rights, Animal Rights, Consumer Electronics, Software Development, Hardware Development, Data Science, Cybersecurity, Astronomy, Chemistry, Physics, Biology, Mathematics, Engineering, Statistics, Linguistics, Archaeology, Anthropology, Genealogy, Gender Studies, Cultural Studies, Media Studies, Communication, Publishing, Design"

Which of the following best describes the content of this page? Provide only one answer with no explanation or additional text: FAQs, Press release, Announcement, Contest, E-book, Landing page, Infographic, White Paper, “How to” post, Blog post, E-course, List Post, Reader question, Roundup, Book review, Survey, Quiz, Photo Album, Video, Case study, General Website Content, Guest Blog or Post, Product Review, Prediction, Demo, Live stream, Awards, Guide, Templates, Check list, Meme, User-Generated Content, Research/Data, Essay, Event Calendar, Giveaway, Podcast, Testimonial, Industry news, Glossary, Comparison Page, Best practices, Recipes, Cheat sheet, Screenshots, Tool or Calculator.

Depth: Identify the depth of coverage on the topic as follows, choosing only one answer from: 1. Topical Coverage (Broad), 2. Module, or 3. Unit, based on where it would fall in the context of a course hierarchy.

Course: If a comprehensive course were to be created about the underlying topic, provide a list of the course modules in chronological order. If the content of this page is considered to be more of a sub-topic, module or unit to a broader course, rather than the "top level" topic, outline broader course.

Experience Level: Choosing only one answer, which level of knowledge/experience is this content intended for? Beginner, Intermediate, Advanced, or None/Not Applicable.

You are generating the output of a JSON file. 

{{input}}
```