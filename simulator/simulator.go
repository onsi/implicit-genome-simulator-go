package simulator

type Simulator struct {
	ImplicitGenome *ImplicitGenome
	Organisms []*Organism
	Environment *Environment
	Time int
}

func NewSimulator(numLoci, numOrganisms int) *Simulator{
	igenome := NewImplicitGenome(numLoci)
	rec := Simulator{
		ImplicitGenome: igenome,
		Environment: NewEnvironment(igenome),
		Organisms: make([]*Organism, numOrganisms),
		Time: 0,
	}

	for i := 0; i < numOrganisms; i++ {
		rec.Organisms[i] = NewOrganism(igenome)
	}

	return &rec
}

func (rec *Simulator) PerformIteration() {
	DataLog(ITERATION_START, nil)
	rec.Time += 1
	Log("Iteration %d: Organisms %d\n", rec.Time, len(rec.Organisms))
	newOrganisms := []*Organism{}
	for _, o := range(rec.Organisms) {
		offspring := o.OffspringForEnvironment(rec.Environment)
		for _, newO := range(offspring) {
			newOrganisms = append(newOrganisms, newO)
		}
	}
	rec.Organisms = newOrganisms

	rec.ChangeEnvironment()

	DataLog(ITERATION_COMPLETE, nil)
}

func (rec *Simulator) PerformIterations(numIterations int) {
	for i := 0; i < numIterations; i++ {
		rec.PerformIteration()
	}
}

func (rec *Simulator) ChangeEnvironment() {
	// FIXME
}

