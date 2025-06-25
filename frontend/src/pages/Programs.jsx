import React, { useEffect, useState } from 'react';

const Programs = () => {
  const [programs, setPrograms] = useState([]);

  useEffect(() => {
  setPrograms([
    {
      id: 1,
      title: "Стипендия Bolashaq",
      country: "Казахстан",
      deadline: "2025-07-31",
      description: "Государственная программа финансирования обучения за рубежом.",
    },
    {
      id: 2,
      title: "Грант Erasmus+",
      country: "Европейский Союз",
      deadline: "2025-10-01",
      description: "Программа академического обмена для студентов и магистров.",
    },
    {
      id: 3,
      title: "Стажировка в OECD",
      country: "Франция",
      deadline: "2025-08-20",
      description: "Оплачиваемая стажировка для молодых специалистов и выпускников.",
    },
    {
      id: 4,
      title: "Грант DAAD",
      country: "Германия",
      deadline: "2025-09-15",
      description: "Финансирование магистратуры и аспирантуры в университетах Германии.",
    },
  ]);
}, []);


  return (
    <div>
      <h1>Все программы 🎓</h1>
      {programs.length === 0 ? (
        <p>Загрузка...</p>
      ) : (
        <ul>
          {programs.map((p) => (
            <div className="program-card" key={p.id}>
              <h3>{p.title}</h3>
              <p><strong>Страна:</strong> {p.country}</p>
              <p><strong>Дедлайн:</strong> {p.deadline}</p>
              <p>{p.description}</p>
              </div>
          ))}
        </ul>
      )}
    </div>
  );
};

export default Programs;
