import React, { useEffect, useRef } from 'react';
import './List.css';
import ListItem from './ListItem';

function List({ data, onItemSelected, highlightItems }) {
  const listRef = useRef(null);

  useEffect(() => {
    if (listRef.current) {
      listRef.current.scrollTop = 0;
    }
  }, [data]);

  const highlightList = highlightItems ? data.filter((i) => i.highlight) : [];
  const list = highlightItems ? data.filter((i) => !i.highlight) : data;

  return (
    <div className="List" ref={listRef}>
      {highlightList.map((itemData) => (
        <ListItem
          highlightItem={true}
          key={itemData.url}
          data={itemData}
          onClick={onItemSelected}
        />
      ))}
      {list.map((itemData) => (
        <ListItem key={itemData.url} data={itemData} onClick={onItemSelected} />
      ))}
    </div>
  );
}

export default List;
