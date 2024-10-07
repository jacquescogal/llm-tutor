
const TabGroup = (props: {
    chosenTab: string;
    tabs: string[];
    setChosenTab: (tabName: string) => void;
  }) => {
    return (
      <div role="tablist" className="tabs tabs-boxed">
          {
              props.tabs.map((tab, index) => (
                  <a
                      key={index}
                      role="tab"
                      className={`tab ${props.chosenTab === tab ? "tab-active" : ""}`}
                      onClick={() => props.setChosenTab(tab)}
                  >
                      {tab}
                  </a>
              ))
          }
      </div>
    );
  };

  export default TabGroup;