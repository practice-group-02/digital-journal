INSERT INTO users (username, email, password_hash, role)
VALUES
    ('user3', 'user3@example.com', '$2a$10$nf0wqX.GQk//MlI5RJXWy.RegNNryZDa5CrAPsGLTS4UvlLdTKEda', 'user');


INSERT INTO program_types (name) VALUES
  ('стипендия'),
  ('грант'),
  ('программа'),
  ('конкурс'),
  ('стажировка');



INSERT INTO tags (name) VALUES
  ('казахстан'),
  ('бакалавриат'),
  ('магистратура'),
  ('phd'),
  ('научные исследования'),
  ('стажировка'),
  ('обучение за рубежом'),
  ('гуманитарные науки'),
  ('технические специальности'),
  ('европа'),
  ('япония'),
  ('венгрия'),
  ('сша'),
  ('гендерное равенство'),
  ('искусство'),
  ('гражданское общество'),
  ('австрия'),
  ('ec'),
  ('международный'),
  ('архитектура'),
  ('инженерия'),
  ('стартап'),
  ('технологии'),
  ('nasa'),
  ('экономика'),
  ('франция'),
  ('турция');





INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'Стипендия «Болашақ»',
    'Государственная стипендиальная программа Республики Казахстан, предоставляющая возможность получения образования за рубежом. Покрывает обучение, проживание, перелёты, медицинскую страховку и другие расходы. Подходит для бакалавриата, магистратуры, PhD, стажировок и научных исследований.',
    1,
    'Казахстан',
    'Центр международных программ при Министерстве науки и высшего образования РК',
    '1900-01-01',
    'https://bolashak.gov.kz/ru/',
    1
    );

INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'Stipendium Hungaricum',
    'Международная стипендиальная программа, финансируемая правительством Венгрии. Предоставляет полную стипендию для иностранных студентов на обучение в венгерских университетах по программам бакалавриата, магистратуры, докторантуры и подготовительного года. Покрывает обучение, медицинскую страховку, ежемесячную стипендию, общежитие или пособие на жильё.',
    1,
    'Венгрия',
    'Tempus Public Foundation',
    '2025-01-15',
    'https://stipendiumhungaricum.hu/',
    1
    );

INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'Стипендия MEXT (Monbukagakusho)',
    'Полностью финансируемая стипендия от правительства Японии для иностранных студентов. Подходит для обучения по программам бакалавриата, магистратуры, аспирантуры, а также для исследований и языковых курсов. Покрывает обучение, авиаперелёты, ежемесячную стипендию и медицинскую страховку.',
    1,
    'Япония',
    'MEXT',
    '1900-01-01',
    'https://www.studyinjapan.go.jp/en/smap-stopj-applications-scholarships-mext/',
    1
    );


INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'Benjamin A. Gilman International Scholarship',
    'Стипендия для студентов с ограниченными финансовыми возможностями на обучение или стажировку за рубежом.',
    1,
    'США',
    'Госдепартамент США',
    '2025-03-06',
    'https://www.gilmanscholarship.org/',
    1
    );


INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'AAUW International Project Grants',
    'Гранты женщинам для реализации проектов по гендерному равенству в своих странах.',
    2,
    'международный',
    'AAUW',
    '1900-01-01',
    'https://www.aauw.org/',
    1
    );

INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'Allianz Foundation Grants Program 2025',
    'Гранты в области искусства, климата и гражданского общества.',
    2,
    'Европа / международный',
    'Allianz Foundation',
    '1900-01-01',
    'https://www.allianzfoundation.org/',
    1
    );

INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'Гранты на обучение в Австрии (Венский университет и Технический университет Граца)',
    'Австрийские вузы, такие как Венский университет и Технический университет Граца, предлагают различные гранты и стипендии для иностранных студентов, включая граждан Казахстана. Доступны программы на бакалавриат, магистратуру и PhD, а также исследовательские гранты. Часто покрывают обучение, проживание и дают ежемесячную стипендию. Возможны гранты от самих университетов, правительства Австрии и организаций типа OeAD.',
    2,
    'Австрия',
    'Австрийское агентство по международной мобильности и сотрудничеству в сфере образования и науки (OeAD), Венский университет, Технический университет Граца',
    '1900-01-01',
    'https://www.univie.ac.at/en/studying/',
    1
    );

INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'Marie Skłodowska‑Curie Actions',
    'Стипендии и гранты ЕС для исследователей в рамках Horizon Europe.',
    3,
    'ЕС / международный',
    'Европейская комиссия',
    '1900-01-01',
    'https://ec.europa.eu/',
    1
    );

INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'Fulbright Scholar Program',
    'Престижная стипендиальная программа правительства США для учёных, преподавателей, исследователей и специалистов. Предоставляет возможность проведения исследований, преподавания или повышения квалификации в университетах и научных учреждениях США. Покрывает все основные расходы: авиаперелёт, проживание, медицинскую страховку, визовую поддержку и ежемесячную стипендию.',
    3,
    'США',
    'Bureau of Educational and Cultural Affairs',
    '2025-05-15',
    'https://kz.usembassy.gov/education-culture/fulbright-programs/',
    1
    );

INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'Erasmus Mundus Joint Master Degrees',
    'Программа Европейского Союза, предоставляющая полные стипендии для обучения по совместным магистерским программам в университетах Европы (и иногда за её пределами). Студенты обучаются в нескольких странах, получают международный диплом и опыт. Покрывает обучение, проживание, проезд и выдаёт ежемесячную стипендию.',
    3,
    'Европейские страны',
    'EACEA — EU Commission',
    '2025-01-10',
    'https://erasmus-plus.ec.europa.eu/opportunities/individuals/students/erasmus-mundus-joint-masters',
    1
    );

INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'Global Undergraduate Exchange Program (Global UGRAD)',
    'Программа академического обмена для студентов бакалавриата, предлагающая бесплатное семестровое обучение в университетах США. Покрывает обучение, проживание, питание, перелёт и страховку. Обязательное условие — возвращение на родину после завершения программы.',
    3,
    'США',
    'Госдепартамент США',
    '2025-01-15',
    'https://www.worldlearning.org/program/global-undergraduate-exchange-program/',
    1
    );

INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'Türkiye Scholarships',
    'Государственная программа стипендий в Турции для бакалавриата, магистратуры и PhD.',
    3,
    'Турция',
    'Правительство Турции',
    '1900-01-01',
    'https://www.turkiyeburslari.gov.tr/',
    1
    );

INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'AEI International Student Design Competition',
    'Междисциплинарный конкурс для студентов архитектуры и инженерии.',
    4,
    'США',
    'Architectural Engineering Institute (AEI)',
    '2025-02-28',
    'https://www.aeisdc.org/home',
    1
    );

INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'Red Bull Basement',
    'Международный конкурс для студенческих стартапов в сфере технологий.',
    4,
    'глобальный',
    'Red Bull',
    '1900-01-01',
    'https://www.redbullbasement.com/',
    1
    );

INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'NASA Internship Program',
    'Стажировка в NASA для студентов и выпускников технических специальностей.',
    5,
    'США',
    'NASA',
    '2025-06-26',
    'https://www.nasa.gov/',
    1
    );

INSERT INTO programs
    (title, description, type, country, organization, deadline, link, user_id)
    VALUES (
    'OECD Internship Programme',
    'Оплачиваемая стажировка в Париже для студентов бакалавриата, магистратуры и PhD.',
    5,
    'Франция',
    'OECD',
    '1900-01-01',
    'https://www.oecd.org/careers/internship-programme/',
    1
    );



-- Стипендия «Болашақ» (ID 1)
INSERT INTO programs_tags VALUES
(1, 1), (1, 2), (1, 3), (1, 4), (1, 5), (1, 6), (1, 7);

-- Stipendium Hungaricum (ID 2)
INSERT INTO programs_tags VALUES
(2, 2), (2, 3), (2, 4), (2, 6), (2, 7), (2, 11);

-- MEXT (ID 3)
INSERT INTO programs_tags VALUES
(3, 2), (3, 3), (3, 4), (3, 5), (3, 7), (3, 10);

-- Gilman Scholarship (ID 4)
INSERT INTO programs_tags VALUES
(4, 6), (4, 7), (4, 13);

-- AAUW Grants (ID 5)
INSERT INTO programs_tags VALUES
(5, 13), (5, 19);

-- Allianz Foundation (ID 6)
INSERT INTO programs_tags VALUES
(6, 14), (6, 15), (6, 16), (6, 19);

-- Grants in Austria (ID 7)
INSERT INTO programs_tags VALUES
(7, 2), (7, 3), (7, 4), (7, 5), (7, 7), (7, 17);

-- Marie Skłodowska‑Curie (ID 8)
INSERT INTO programs_tags VALUES
(8, 4), (8, 5), (8, 18), (8, 19);

-- Fulbright (ID 9)
INSERT INTO programs_tags VALUES
(9, 5), (9, 6), (9, 7), (9, 13);

-- Erasmus Mundus (ID 10)
INSERT INTO programs_tags VALUES
(10, 3), (10, 7), (10, 10), (10, 19);

-- Global UGRAD (ID 11)
INSERT INTO programs_tags VALUES
(11, 2), (11, 6), (11, 13);

-- Türkiye Scholarships (ID 12)
INSERT INTO programs_tags VALUES
(12, 2), (12, 3), (12, 4), (12, 26);

-- AEI Design Competition (ID 13)
INSERT INTO programs_tags VALUES
(13, 20), (13, 21), (13, 19);

-- Red Bull Basement (ID 14)
INSERT INTO programs_tags VALUES
(14, 22), (14, 23), (14, 19);

-- NASA Internship (ID 15)
INSERT INTO programs_tags VALUES
(15, 6), (15, 9), (15, 24);

-- OECD Internship (ID 16)
INSERT INTO programs_tags VALUES
(16, 6), (16, 25), (16, 19);
