package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Patient struct {
	PatientID     string
	Gender        string
	Age           int
	Neighbourhood string
	Scholarship   bool
	Hipertension  bool
	Diabetes      bool
	Alcoholism    bool
	Handcap       bool
	VisitDate     string
}

type Appointment struct {
	AppointmentID   string
	PatientID       string
	ScheduledDay    string
	AppointmentDay  string
	AppointmentDate string
	NoShow          bool
}

type Clinic struct {
	ClinicID      string
	Neighbourhood string
	Capacity      int
}

type Doctor struct {
	DoctorID     string
	DoctorName   string
	Specialty    string
	ClinicID     string
	Availability string
	DoctorDate   string
}

type LastReport struct {
	PatientID      string
	DoctorID       string
	LastReportDate string
	HealthDisease  string
	LastBilling    float64
}

var neighbourhoodNames = []string{
	"Greenwood", "Sunnydale", "Riverside", "Brooklyn", "Oakwood",
	"Mapleton", "Silverstone", "Pinehill", "Lakeside", "Eastwood",
	"Brighton", "Shady Grove", "Westfield", "Clearwater", "Cedar Park",
}

func getRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func randomBoolean() bool {
	return rand.Intn(2) == 0
}

func selectString(probabilities map[string]float64) string {
	total := 0.0
	for _, p := range probabilities {
		total += p
	}

	randVal := rand.Float64() * total
	cumulative := 0.0
	for key, prob := range probabilities {
		cumulative += prob
		if randVal <= cumulative {
			return key
		}
	}
	return ""
}

func writeToCSV(filename string, data [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func generateID(prefix string, index int) string {
	return fmt.Sprintf("%s%04d", prefix, index)
}

func generateData() {
	// Generate Patients data
	patients := []Patient{}
	for i := 1; i <= 100; i++ {
		patient := Patient{
			PatientID:     generateID("P", i),
			Gender:        selectString(map[string]float64{"Male": 0.5, "Female": 0.5}),
			Age:           getRandomNumber(18, 90),
			Neighbourhood: neighbourhoodNames[rand.Intn(len(neighbourhoodNames))],
			Scholarship:   randomBoolean(),
			Hipertension:  randomBoolean(),
			Diabetes:      randomBoolean(),
			Alcoholism:    randomBoolean(),
			Handcap:       randomBoolean(),
			VisitDate:     time.Now().Add(time.Duration(-rand.Intn(30)) * time.Hour * 24).Format("2006-01-02"),
		}
		patients = append(patients, patient)
	}

	// Write Patients data to CSV
	patientData := [][]string{{"PatientID", "Gender", "Age", "Neighbourhood", "Scholarship", "Hipertension", "Diabetes", "Alcoholism", "Handcap", "VisitDate"}}
	for _, p := range patients {
		patientData = append(patientData, []string{
			p.PatientID, p.Gender, strconv.Itoa(p.Age), p.Neighbourhood, strconv.FormatBool(p.Scholarship),
			strconv.FormatBool(p.Hipertension), strconv.FormatBool(p.Diabetes), strconv.FormatBool(p.Alcoholism), strconv.FormatBool(p.Handcap),
			p.VisitDate,
		})
	}
	writeToCSV("patients.csv", patientData)

	// Generate Appointments data
	appointments := []Appointment{}
	for i, patient := range patients {
		appointment := Appointment{
			AppointmentID:   generateID("A", i+1),
			PatientID:       patient.PatientID,
			ScheduledDay:    time.Now().Add(time.Duration(-rand.Intn(30)) * time.Hour * 24).Format("2006-01-02"),
			AppointmentDay:  time.Now().Add(time.Duration(rand.Intn(30)) * time.Hour * 24).Format("2006-01-02"),
			AppointmentDate: time.Now().Add(time.Duration(-rand.Intn(15)) * time.Hour * 24).Format("2006-01-02"),
			NoShow:          randomBoolean(),
		}
		appointments = append(appointments, appointment)
	}

	// Write Appointments data to CSV
	appointmentData := [][]string{{"AppointmentID", "PatientID", "ScheduledDay", "AppointmentDay", "AppointmentDate", "NoShow"}}
	for _, a := range appointments {
		appointmentData = append(appointmentData, []string{
			a.AppointmentID, a.PatientID, a.ScheduledDay, a.AppointmentDay, a.AppointmentDate, strconv.FormatBool(a.NoShow),
		})
	}
	writeToCSV("appointments.csv", appointmentData)

	// Generate Clinics data
	clinics := []Clinic{}
	for i := 1; i <= 10; i++ {
		clinic := Clinic{
			ClinicID:      generateID("C", i),
			Neighbourhood: neighbourhoodNames[rand.Intn(len(neighbourhoodNames))],
			Capacity:      getRandomNumber(5, 20),
		}
		clinics = append(clinics, clinic)
	}

	// Write Clinics data to CSV
	clinicData := [][]string{{"ClinicID", "Neighbourhood", "Capacity"}}
	for _, c := range clinics {
		clinicData = append(clinicData, []string{
			c.ClinicID, c.Neighbourhood, strconv.Itoa(c.Capacity),
		})
	}
	writeToCSV("clinics.csv", clinicData)

	// Generate Doctors data
	doctors := []Doctor{}
	specialties := []string{"Cardiologist", "Orthopedist", "Dermatologist", "Neurologist", "Pediatrician", "Psychiatrist"}
	for i := 1; i <= 20; i++ {
		doctor := Doctor{
			DoctorID:     generateID("D", i),
			DoctorName:   fmt.Sprintf("Doctor %d", i),
			Specialty:    specialties[rand.Intn(len(specialties))],
			ClinicID:     generateID("C", rand.Intn(len(clinics))+1),
			Availability: fmt.Sprintf("%02d:%02d-%02d:%02d", rand.Intn(8)+8, rand.Intn(60), rand.Intn(8)+12, rand.Intn(60)),
			DoctorDate:   time.Now().Add(time.Duration(-rand.Intn(30)) * time.Hour * 24).Format("2006-01-02"),
		}
		doctors = append(doctors, doctor)
	}

	// Write Doctors data to CSV
	doctorData := [][]string{{"DoctorID", "DoctorName", "Specialty", "ClinicID", "Availability", "DoctorDate"}}
	for _, d := range doctors {
		doctorData = append(doctorData, []string{
			d.DoctorID, d.DoctorName, d.Specialty, d.ClinicID, d.Availability, d.DoctorDate,
		})
	}
	writeToCSV("doctors.csv", doctorData)

	fmt.Println("Data generation complete.")
}

func generateLastReports() {
	lastReports := []LastReport{}
	for i := 1; i <= 100; i++ {
		patientID := generateID("P", i)
		healthConditions := []string{}

		if randomBoolean() {
			healthConditions = append(healthConditions, "Hipertension")
		}
		if randomBoolean() {
			healthConditions = append(healthConditions, "Diabetes")
		}
		if randomBoolean() {
			healthConditions = append(healthConditions, "Alcoholism")
		}
		if randomBoolean() {
			healthConditions = append(healthConditions, "Handcap")
		}

		lastReport := LastReport{
			PatientID:      patientID,
			DoctorID:       generateID("D", rand.Intn(10)+1),
			LastReportDate: time.Now().Add(time.Duration(-rand.Intn(90)) * time.Hour * 24).Format("2006-01-02"),
			HealthDisease:  strings.Join(healthConditions, ", "),
			LastBilling:    float64(getRandomNumber(50, 500)),
		}

		lastReports = append(lastReports, lastReport)
	}

	lastReportData := [][]string{{"PatientID", "DoctorID", "LastReportDate", "HealthDisease", "LastBilling"}}
	for _, lr := range lastReports {
		lastReportData = append(lastReportData, []string{
			lr.PatientID, lr.DoctorID, lr.LastReportDate, lr.HealthDisease, fmt.Sprintf("%.2f", lr.LastBilling),
		})
	}
	writeToCSV("last_reports.csv", lastReportData)

	fmt.Println("Last reports generated.")
}

func main() {
	generatePatients := flag.Bool("generate-patients", false, "Generate patient records")
	generateLastReportsFlag := flag.Bool("generate-last-reports", false, "Generate last report records")
	generateAll := flag.Bool("generate-all", false, "Generate all records")
	flag.Parse()

	if *generatePatients {
		generateData()
	} else if *generateLastReportsFlag {
		generateLastReports()
	} else if *generateAll {
		generateData()
		generateLastReports()
	} else {
		fmt.Println("Please provide a valid flag: -generate-patients, -generate-last-reports, or -generate-all")
	}
}
