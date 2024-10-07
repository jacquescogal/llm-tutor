import { useDispatch } from "react-redux";
import { clearLinkList, popLink, pushLink, setLinks, setRootLink } from "../breadcrumbSlice";
import { useLocation } from 'react-router-dom';
import { useSelector } from "react-redux";
import { RootState } from "../store";

type Params = {
  name: string;
};

export const CrumbHelperSetRootLink = (params: Params) => {
  const dispatch = useDispatch();
  const location = useLocation();
  dispatch(setRootLink({
    name: params.name,
    path: location.pathname
  }));
  return null;
};

export const CrumbHelperPushLink = (params: Params) => {
  const dispatch = useDispatch();
  const location = useLocation();
  const subsetFound = CrumbHelperFindLinkSubset(params.name);
  const lastLink = CrumbHelperPeekLink();
  if (lastLink !== null && lastLink.name === params.name) {
    return null;
  }
  else if (subsetFound !== null) {
    dispatch(setLinks(subsetFound));
    return null;
  }
  dispatch(pushLink({
    name: params.name,
    path: location.pathname
  }),[]);
  return null;
};

export const CrumbHelperPopLink = () => {
  const dispatch = useDispatch();
  dispatch(popLink());
  return null;
}

export const CrumbHelperClearLinkList = () => {
  const dispatch = useDispatch();
  dispatch(clearLinkList());
  return null;
}

export const CrumbHelperPeekLink = () => {
  const breadCrumbList = useSelector((state: RootState) => state.breadCrumbList);
  if (breadCrumbList.value.length === 0) {
    return null;
  }
  return breadCrumbList.value[breadCrumbList.value.length - 1];
}

export const CrumbHelperFindLinkSubset= (name: string) => {
  const breadCrumbList = useSelector((state: RootState) => state.breadCrumbList);
  const lastIndex = breadCrumbList.value.findIndex((crumb) => crumb.name === name);
  if (lastIndex === -1) {
    return null;
  }
  return breadCrumbList.value.slice(0, lastIndex + 1);
}