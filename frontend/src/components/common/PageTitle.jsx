function PageTitle({ title, description = "" }) {
  return (
    <section className="page-title">
      <h2>{title}</h2>
      {description ? <p>{description}</p> : null}
    </section>
  );
}

export default PageTitle;
