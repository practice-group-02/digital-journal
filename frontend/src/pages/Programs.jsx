import React, { useEffect, useState } from 'react';

const Programs = () => {
  const [programs, setPrograms] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/programs')
      .then((res) => res.json())
      .then((data) => setPrograms(data))
      .catch((err) => console.error('Ошибка при получении программ:', err));
  }, []);

  return (
    <div>
      <h1>Все программы 🎓</h1>
      {programs.length === 0 ? (
        <p>Загрузка...</p>
      ) : (
        <ul>
          {programs.map((p) => (
            <li key={p.id}>
              <h3>{p.title}</h3>
              <p>{p.country}</p>
              <p>Дедлайн: {p.deadline}</p>
              <p>{p.description}</p>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default Programs;
