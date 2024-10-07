import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import LoginPage from "./pages/LoginPage";
import BasePage from "./pages/BasePage";
import ExploreSubjectPage from "./pages/explore/ExploreSubjectPage";
import ExploreModulePage from "./pages/explore/ExploreModulePage";
import SubjectPage from "./pages/focus/SubjectPage";
import ModulePage from "./pages/focus/ModulePage";
import DocumentPage from "./pages/focus/DocumentPage";
import QuestionPage from "./pages/questions/QuestionPage";
import FavouriteSubjectPage from "./pages/favourites/FavouriteSubjectPage";
import FavouriteModulePage from "./pages/favourites/FavouriteModulePage";
function App() {
  // // access state with useSelector, typed with RootState
  // const count = useSelector((state: RootState) => state);

  // // Use AppDispatch type for dispatch
  // const dispatch = useDispatch();
  return (
    <>
      {/* <div>
        {count.counter.value}
        <button onClick={() => dispatch(increment())}>+</button>
        <button onClick={() => dispatch(decrement())}>-</button>
        <button onClick={() => dispatch(incrementByAmount(2))}>+2</button>
      </div> */}
      <div className="fixed left-0 top-0 h-screen w-screen bg-base-100">
        <Router>
          {/* Header banner */}
          {/* <Header userSession={userSession} endSession={endSession}/> */}
          {/* Routes to pages */}
          <Routes>
            {/* <Route path="/" element={<Navigate to="/login" replace />} /> */}
            <Route path="/dashboard" element={<BasePage />}>
              <Route path="login" element={<LoginPage />} />
              <Route path="explore">
                <Route path="subject" element={<ExploreSubjectPage />} />
                <Route path="subject/:subjectId" element={<SubjectPage />} />
                <Route path="subject/:subjectId/module/:moduleId" element={<ModulePage />} />
                <Route path="subject/:subjectId/module/:moduleId/document/:documentId" element={<DocumentPage />} />
                <Route path="module" element={<ExploreModulePage />} />
                <Route path="module/:moduleId" element={<ModulePage />} />
                <Route path="module/:moduleId/document/:documentId" element={<DocumentPage />} />
              </Route>
              <Route path="favourite">
                <Route path="subject" element={<FavouriteSubjectPage />} />
                <Route path="subject/:subjectId" element={<SubjectPage />} />
                <Route path="subject/:subjectId/module/:moduleId" element={<ModulePage />} />
                <Route path="subject/:subjectId/module/:moduleId/document/:documentId" element={<DocumentPage />} />
                <Route path="module" element={<FavouriteModulePage />} />
                <Route path="module/:moduleId" element={<ModulePage />} />
                <Route path="module/:moduleId/document/:documentId" element={<DocumentPage />} />
              </Route>
              <Route path="question" element={<QuestionPage/>}/>
            </Route>
            <Route path="subject/:subjectId" element={<SubjectPage />} />
            <Route path="/login" element={<LoginPage />} />
          </Routes>
        </Router>
      </div>
    </>
  );
}

export default App;
